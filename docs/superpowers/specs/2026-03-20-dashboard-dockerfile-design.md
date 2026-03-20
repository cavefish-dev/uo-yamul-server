# dashboard.Dockerfile — Design Spec

**Date:** 2026-03-20
**Branch:** local-dev-improvements
**Status:** Approved

## Goal

Add a `dashboard.Dockerfile` that builds the Flutter web dashboard and serves it via nginx. Update `docker-compose.yml` so the full mono-repo can be run locally with a single `docker-compose up`, exposing only the gateway (TCP :2593) and the dashboard (HTTP :8080) externally.

## Architecture

```
Browser (localhost:8080)
    │
    ▼
yamul-dashboard container (nginx :8080) ── EXTERNAL
    ├── GET /*        → serve Flutter web static files
    └── POST /grpc/*  → proxy_pass → yamul-envoy:8091 (internal)
            │
            ▼ (internal docker network)
yamul-envoy container (Envoy :8091) ── INTERNAL ONLY
    └── gRPC-Web → gRPC translation → yamul-backend:8087

yamul-backend (Kotlin/Spring Boot) ── INTERNAL ONLY
    ├── :8087  Login gRPC   ← wired to Envoy
    ├── :8088  Character gRPC
    └── :8089  Game gRPC

UO Client (TCP :2593)
    ▼
yamul-gateway (Go) ── EXTERNAL
    └── gRPC → yamul-backend (internal)
```

**Scope note:** Envoy is wired only to `:8087` (Login service) because the dashboard's current gRPC client (`DashboardLoginServiceClient`) only calls the Login service. Character and Game services are not proxied via gRPC-Web in this phase.

**Known limitation:** `yamul-gateway` has no `depends_on: yamul-backend` in the current compose file and this spec does not add one (pre-existing gap, separate concern from dashboard work).

## Components

### 1. `dashboard.Dockerfile`

Two-stage build. Build context is repo root (consistent with existing Dockerfiles).

**Stage 1 — `flutter-builder`**

**Prerequisite fix:** `yamul-dashboard/pubspec.yaml` line 10 currently reads `sdk: 3.11.3` (bare version, interpreted as exact match `=3.11.3`). This must be changed to a range constraint (`'^3.11.3'` or `'>=3.11.3 <4.0.0'`) before implementation so that compatible Flutter images can be used. This fix is a prerequisite, not part of this feature per se, but the Docker build will fail without it.

Base image: `ghcr.io/cirruslabs/flutter` at a tag whose bundled Dart SDK satisfies the `sdk` constraint in `pubspec.yaml`. Verify the tag against the [cirruslabs Flutter release list](https://github.com/cirruslabs/docker-images-flutter/releases) before implementation.

Stage 1 `WORKDIR` is `/app` (the Flutter project root). All paths below are relative to `/app`.

Steps:
1. Copy `api-definitions/dashboard/proto/` → `/app/proto/`
2. Install `protoc` and activate `protoc-gen-dart` via `dart pub global activate protoc_plugin` (mirrors what `devBuild.ps1` does)
3. Run proto generation using absolute paths (consistent with `devBuild.ps1`):
   ```
   protoc \
     --plugin=protoc-gen-dart=/root/.pub-cache/bin/protoc-gen-dart \
     --proto_path=/app/proto \
     --dart_out=grpc:/app/lib/generated/grpc \
     /app/proto/*.proto
   ```
4. Copy `pubspec.yaml` + `pubspec.lock`, run `flutter pub get`
5. Copy remaining Flutter source, run `flutter build web --release`
6. Built assets are at `/app/build/web`

The generated stubs (`lib/generated/grpc/`) are not committed to the repo, so the Docker build must produce them.

**Stage 2 — `nginx:alpine`**

- Copy built web assets from Stage 1: `--from=flutter-builder /app/build/web /usr/share/nginx/html`
- Copy nginx config: `COPY ./yamul-dashboard/nginx.conf /etc/nginx/conf.d/default.conf`
- `EXPOSE 8080`

### 2. `yamul-dashboard/nginx.conf`

```nginx
server {
    listen 8080;

    location / {
        root /usr/share/nginx/html;
        index index.html;
        try_files $uri $uri/ /index.html;
    }

    location /grpc/ {
        proxy_pass         http://yamul-envoy:8091/;
        proxy_http_version 1.1;
        proxy_buffering    off;
        proxy_set_header   Host         $host;
        proxy_set_header   Content-Type $http_content_type;
    }
}
```

`proxy_buffering off` is required so gRPC-Web streaming responses are forwarded immediately rather than buffered. `proxy_http_version 1.1` enables keep-alive to Envoy.

### 3. Envoy config (`api-definitions/dashboard/envoy.yaml`)

Two changes:
- `virtual_hosts[0].name`: rename from `wiremock` → `yamul-backend`
- Cluster `load_assignment` endpoint: change `address: wiremock`, `port_value: 8080` → `address: yamul-backend`, `port_value: 8087`

Everything else (gRPC-Web filter, CORS headers, listener port 8091) stays unchanged.

### 4. `docker-compose.yml`

| Service | Change |
|---|---|
| `yamul-gateway` | Unchanged — port 2593 external |
| `yamul-backend` | Remove all external port mappings (8087/8088/8089 internal only); add `healthcheck` (e.g., `grpc_health_probe` or TCP check on :8087) |
| `yamul-envoy` | New — `envoyproxy/envoy:v1.31-latest`; volume mount `./api-definitions/dashboard/envoy.yaml:/etc/envoy/envoy.yaml`; entrypoint: `/usr/local/bin/envoy -c /etc/envoy/envoy.yaml --service-cluster proxy`; no external ports; `depends_on: yamul-backend: condition: service_healthy` |
| `yamul-dashboard` | New — built from `dashboard.Dockerfile`; port 8080 external; `depends_on: yamul-envoy: condition: service_started` |

All services share the default docker-compose network and communicate by service name.

**Backend healthcheck:** Use a bash TCP check (works on the Corretto/Amazon Linux base without installing `netcat`):
```yaml
healthcheck:
  test: ["CMD-SHELL", "bash -c 'echo > /dev/tcp/localhost/8087'"]
  interval: 10s
  timeout: 5s
  retries: 10
  start_period: 60s
```
`start_period: 60s` accounts for the Gradle/JVM startup time.

### 5. `yamul-dashboard/lib/core/grpc/login_grpc_service.dart`

Change the hardcoded `http://localhost:8091` endpoint to derive the base URL from the current page origin at runtime:

```dart
final uri = Uri(
  scheme: Uri.base.scheme,
  host: Uri.base.host,
  port: Uri.base.port,
  path: '/grpc/',
);
final channel = GrpcWebClientChannel.xhr(uri);
```

`Uri.base` is a Flutter web API (resolves to the browser's current page URL). This file is only used in the web build — it does not affect Android/Linux/Windows targets.

## Files Changed

| File | Action |
|---|---|
| `yamul-dashboard/pubspec.yaml` | Edit — prerequisite fix: change `sdk: 3.11.3` → `sdk: '^3.11.3'` |
| `dashboard.Dockerfile` | Create |
| `yamul-dashboard/nginx.conf` | Create |
| `api-definitions/dashboard/envoy.yaml` | Edit — rename virtual host, change upstream to `yamul-backend:8087` |
| `docker-compose.yml` | Edit — add dashboard + envoy services, remove backend external ports, add backend healthcheck |
| `yamul-dashboard/lib/core/grpc/login_grpc_service.dart` | Edit — dynamic gRPC URL from `Uri.base` |

## Out of Scope

- Production deployment (TLS, reverse proxy, domain)
- Dashboard authentication or other Dart service files (only `login_grpc_service.dart` is changed)
- Backend gRPC-Web native support
- CI/CD changes
- Wiring Envoy to Character (:8088) or Game (:8089) services
- `yamul-gateway` startup ordering against backend
