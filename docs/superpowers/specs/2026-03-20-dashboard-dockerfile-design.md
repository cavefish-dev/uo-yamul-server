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
    ├── :8087  Login gRPC
    ├── :8088  Character gRPC
    └── :8089  Game gRPC

UO Client (TCP :2593)
    ▼
yamul-gateway (Go) ── EXTERNAL
    └── gRPC → yamul-backend (internal)
```

## Components

### 1. `dashboard.Dockerfile`

Two-stage build:

- **Stage 1 (`flutter-builder`)**: Uses a Flutter SDK Docker image. Copies `pubspec.yaml`, runs `flutter pub get`, copies source, runs `flutter build web --release`.
- **Stage 2 (`nginx:alpine`)**: Copies built web assets from Stage 1 into `/usr/share/nginx/html`. Copies a custom `nginx.conf`. Exposes port 8080.

### 2. `nginx.conf` (bundled in the image)

Single server block on port 8080:

- `location /` — serves Flutter SPA with `try_files $uri $uri/ /index.html` for client-side routing.
- `location /grpc/` — `proxy_pass http://yamul-envoy:8091/` forwarding gRPC-Web requests to the internal Envoy service.

### 3. Envoy config (`api-definitions/dashboard/envoy.yaml`)

The existing file is reused with one change: the upstream cluster endpoint changes from `wiremock:8080` to `yamul-backend:8087`. All other config (gRPC-Web filter, CORS headers, port 8091) remains unchanged.

### 4. `docker-compose.yml`

Changes:

| Service | Change |
|---|---|
| `yamul-gateway` | Unchanged — port 2593 external |
| `yamul-backend` | Remove all external port mappings (8087/8088/8089 become internal only) |
| `yamul-envoy` | New service — `envoyproxy/envoy:v1.31-latest`, mounts `api-definitions/dashboard/envoy.yaml`, no external ports, `depends_on: yamul-backend` |
| `yamul-dashboard` | New service — built from `dashboard.Dockerfile`, port 8080 external, `depends_on: yamul-envoy` |

All services share the default docker-compose network and communicate by service name.

### 5. `yamul-dashboard/lib/core/grpc/login_grpc_service.dart`

Change the hardcoded `http://localhost:8091` endpoint to derive the gRPC base URL from the page origin at runtime:

```dart
final uri = Uri(
  scheme: Uri.base.scheme,
  host: Uri.base.host,
  port: Uri.base.port,
  path: '/grpc/',
);
final channel = GrpcWebClientChannel.xhr(uri);
```

This makes the same build work at any host:port without build-time configuration.

## Files Changed

| File | Action |
|---|---|
| `dashboard.Dockerfile` | Create |
| `nginx.conf` (in repo root or `yamul-dashboard/`) | Create |
| `api-definitions/dashboard/envoy.yaml` | Edit — change upstream from `wiremock:8080` to `yamul-backend:8087` |
| `docker-compose.yml` | Edit — add dashboard + envoy services, remove backend external ports |
| `yamul-dashboard/lib/core/grpc/login_grpc_service.dart` | Edit — dynamic gRPC URL from page origin |

## Out of Scope

- Production deployment (TLS, reverse proxy, domain)
- Dashboard authentication or other Dart service files (only `login_grpc_service.dart` is changed)
- Backend gRPC-Web native support
- CI/CD changes
