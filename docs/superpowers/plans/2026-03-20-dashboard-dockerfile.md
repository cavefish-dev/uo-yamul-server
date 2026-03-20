# Dashboard Dockerfile & Local Dev Compose Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add `dashboard.Dockerfile` + nginx config to containerise the Flutter web dashboard, add Envoy as an internal gRPC-Web proxy, and update `docker-compose.yml` so `docker-compose up` runs the entire mono-repo with only gateway:2593 and dashboard:8080 accessible from the host.

**Architecture:** Two new docker-compose services — `yamul-dashboard` (nginx serving Flutter web + proxying `/grpc/` to Envoy) and `yamul-envoy` (Envoy translating gRPC-Web to gRPC for the backend Login service). Backend ports become internal-only. The Flutter app derives its gRPC endpoint from `Uri.base` at runtime so the same build works at any host.

**Tech Stack:** Flutter web (Dart), nginx:alpine, Envoy proxy v1.31, Docker multi-stage builds, docker-compose v3.8, grpc_web Dart package.

**Spec:** `docs/superpowers/specs/2026-03-20-dashboard-dockerfile-design.md`

---

## File Map

| File | Action | Responsibility |
|---|---|---|
| `yamul-dashboard/pubspec.yaml` | Edit line 10 | Fix SDK constraint from bare version to range |
| `yamul-dashboard/nginx.conf` | Create | nginx server config: SPA file serving + gRPC-Web proxy |
| `dashboard.Dockerfile` | Create | 2-stage: Flutter build → nginx serve |
| `api-definitions/dashboard/envoy.yaml` | Edit lines 21 + 67-68 | Point Envoy at `yamul-backend:8087` instead of WireMock |
| `docker-compose.yml` | Edit | Add dashboard + envoy services, lock down backend ports, add healthcheck |
| `yamul-dashboard/lib/core/grpc/login_grpc_service.dart` | Edit lines 7-9 | Derive gRPC endpoint from `Uri.base` instead of hardcoded URL |

---

## Task 1: Fix pubspec.yaml SDK constraint

**Files:**
- Modify: `yamul-dashboard/pubspec.yaml:10`

The current value `sdk: 3.11.3` is interpreted as an exact-match constraint (`=3.11.3`). Changing it to a caret range allows any compatible Flutter Docker image to satisfy it.

- [ ] **Step 1: Edit pubspec.yaml**

In `yamul-dashboard/pubspec.yaml`, change line 10 from:
```yaml
  sdk: 3.11.3
```
to:
```yaml
  sdk: '^3.11.3'
```

- [ ] **Step 2: Commit**

```bash
git add yamul-dashboard/pubspec.yaml
git commit -m "fix: relax pubspec SDK constraint to range for Docker compatibility"
```

---

## Task 2: Create nginx.conf

**Files:**
- Create: `yamul-dashboard/nginx.conf`

nginx serves the Flutter SPA and reverse-proxies `/grpc/` to the internal Envoy container. `proxy_buffering off` is required for gRPC-Web streaming. The trailing slash on `proxy_pass http://yamul-envoy:8091/` is intentional — nginx strips the `/grpc/` prefix before forwarding, so gRPC paths like `/yamul.dashboard.DashboardLoginService/Login` arrive at Envoy correctly rooted at `/`. Do not remove the trailing slash.

- [ ] **Step 1: Create the file**

Create `yamul-dashboard/nginx.conf` with this exact content:

```nginx
server {
    listen 8080;

    location / {
        root  /usr/share/nginx/html;
        index index.html;
        try_files $uri $uri/ /index.html;
    }

    location /grpc/ {
        # Trailing slash on proxy_pass strips the /grpc/ prefix before forwarding to Envoy.
        # gRPC paths (e.g. /yamul.dashboard.DashboardLoginService/Login) reach Envoy unchanged.
        proxy_pass         http://yamul-envoy:8091/;
        proxy_http_version 1.1;
        proxy_buffering    off;
        proxy_set_header   Host         $host;
        proxy_set_header   Content-Type $http_content_type;
    }
}
```

- [ ] **Step 2: Commit**

```bash
git add yamul-dashboard/nginx.conf
git commit -m "feat: add nginx config for dashboard SPA serving and gRPC-Web proxy"
```

---

## Task 3: Update Envoy config

**Files:**
- Modify: `api-definitions/dashboard/envoy.yaml:21,67-68`

Two surgical edits: rename the virtual host from `wiremock` to `yamul-backend` (cosmetic), and change the upstream cluster endpoint from `wiremock:8080` to `yamul-backend:8087` (functional — this is the Login gRPC service port).

- [ ] **Step 1: Rename virtual host**

In `api-definitions/dashboard/envoy.yaml` line 21, change:
```yaml
                    - name: wiremock
```
to:
```yaml
                    - name: yamul-backend
```

- [ ] **Step 2: Change upstream endpoint**

In `api-definitions/dashboard/envoy.yaml` lines 67-68, change:
```yaml
                      address: wiremock
                      port_value: 8080
```
to:
```yaml
                      address: yamul-backend
                      port_value: 8087
```

- [ ] **Step 3: Commit**

```bash
git add api-definitions/dashboard/envoy.yaml
git commit -m "feat: point Envoy upstream at yamul-backend Login service"
```

---

## Task 4: Create dashboard.Dockerfile

**Files:**
- Create: `dashboard.Dockerfile`

**Prerequisite:** Task 1 must be completed first — `pubspec.yaml` must have the relaxed SDK constraint or `flutter pub get` will reject the Flutter image's Dart version.

Two-stage build. Stage 1 installs tools, copies source, generates Dart proto stubs (not committed to repo), then builds Flutter web. Stage 2 is a minimal nginx:alpine image.

**Before writing the file — find the Flutter image tag:**

The `ghcr.io/cirruslabs/flutter` image is the community-maintained Flutter Docker image (there is no official Flutter image on Docker Hub). Check https://github.com/cirruslabs/docker-images-flutter/releases for a release whose bundled Dart SDK is `>=3.11.3`. As of early 2026, tags around `3.32.x`–`3.33.x` are expected to satisfy this; pick the lowest tag that meets the constraint. Substitute that tag for `<FLUTTER_TAG>` in the Dockerfile below.

The `ghcr.io/cirruslabs/flutter` image is Ubuntu-based, so `apt-get` is available — this is what allows installing `protobuf-compiler`. The `protoc-gen-dart` plugin is installed to `/root/.pub-cache/bin/protoc-gen-dart`, matching `yamul-dashboard/devBuild.ps1`.

**Critical note on build order:** The proto stubs are generated **after** `COPY ./yamul-dashboard/ .` on purpose. The repo contains `yamul-dashboard/lib/generated/grpc/.keep` (a placeholder). If stubs were generated before the source copy, that copy would overwrite them with the `.keep` file and `flutter build web` would fail with missing-import errors.

- [ ] **Step 1: Create dashboard.Dockerfile**

Create `dashboard.Dockerfile` in the repo root (substitute the real tag for `<FLUTTER_TAG>`):

```dockerfile
# ── Stage 1: Build Flutter web ────────────────────────────────────────────────
FROM ghcr.io/cirruslabs/flutter:<FLUTTER_TAG> AS flutter-builder

WORKDIR /app

# Install protoc (Ubuntu-based image, apt-get is available)
# protoc-gen-dart ends up at /root/.pub-cache/bin/protoc-gen-dart (same as devBuild.ps1)
RUN apt-get update -qq \
    && apt-get install -y -qq protobuf-compiler \
    && dart pub global activate protoc_plugin

# Fetch dependencies first (separate layer — only re-runs if pubspec changes)
COPY ./yamul-dashboard/pubspec.yaml ./yamul-dashboard/pubspec.lock ./
RUN flutter pub get

# Copy Flutter source (overwrites lib/generated/grpc/ placeholder with .keep file — intentional)
COPY ./yamul-dashboard/ .

# Copy proto files and generate stubs AFTER source copy.
# Must be after COPY to avoid the source copy overwriting the generated stubs.
COPY ./api-definitions/dashboard/proto/ /app/proto/
RUN protoc \
      --plugin=protoc-gen-dart=/root/.pub-cache/bin/protoc-gen-dart \
      --proto_path=/app/proto \
      --dart_out=grpc:/app/lib/generated/grpc \
      /app/proto/*.proto

RUN flutter build web --release

# ── Stage 2: Serve with nginx ─────────────────────────────────────────────────
FROM nginx:alpine

COPY --from=flutter-builder /app/build/web /usr/share/nginx/html
COPY ./yamul-dashboard/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 8080
```

- [ ] **Step 2: Test the image builds**

Run from the repo root (this may take several minutes on first run):
```bash
docker build -f dashboard.Dockerfile -t yamul-dashboard:test .
```
Expected: build completes successfully. Watch for:
- `apt-get install protobuf-compiler` and `dart pub global activate` succeed
- `flutter pub get` succeeds (if it errors about SDK version mismatch, pick a different Flutter tag)
- `protoc` generates `.dart` files (not just the `.keep` placeholder)
- `flutter build web --release` succeeds
- Final image is based on `nginx:alpine`

- [ ] **Step 3: Smoke-test the image**

```bash
docker run --rm -p 8080:8080 yamul-dashboard:test
```
Open http://localhost:8080 in a browser.
Expected: Flutter app loads (loading spinner or login screen). No 404 on the root path.

Stop the container with Ctrl-C.

- [ ] **Step 4: Commit**

```bash
git add dashboard.Dockerfile
git commit -m "feat: add dashboard.Dockerfile with Flutter build and nginx serve stages"
```

---

## Task 5: Update docker-compose.yml

**Files:**
- Modify: `docker-compose.yml`

Three changes: (1) remove backend external port mappings, (2) add backend healthcheck, (3) add `yamul-envoy` and `yamul-dashboard` services.

- [ ] **Step 1: Rewrite docker-compose.yml**

Replace the entire file with:

```yaml
version: '3.8'

services:
  yamul-gateway:
    build:
      context: .
      dockerfile: ./gateway.Dockerfile
    ports:
      - "2593:2593"

  yamul-backend:
    build:
      context: .
      dockerfile: ./backend.Dockerfile
    healthcheck:
      # Uses bash TCP pseudo-device — works on the Corretto/Amazon Linux base without netcat
      test: ["CMD-SHELL", "bash -c 'echo > /dev/tcp/localhost/8087'"]
      interval: 10s
      timeout: 5s
      retries: 10
      start_period: 60s

  yamul-envoy:
    image: envoyproxy/envoy:v1.31-latest
    entrypoint: /usr/local/bin/envoy -c /etc/envoy/envoy.yaml --service-cluster proxy
    volumes:
      - ./api-definitions/dashboard/envoy.yaml:/etc/envoy/envoy.yaml
    depends_on:
      yamul-backend:
        condition: service_healthy

  yamul-dashboard:
    build:
      context: .
      dockerfile: ./dashboard.Dockerfile
    ports:
      - "8080:8080"
    depends_on:
      - yamul-envoy
```

- [ ] **Step 2: Validate the compose file**

```bash
docker-compose config
```
Expected: prints the resolved config with no errors. Confirm:
- `yamul-backend` has no `ports` section
- `yamul-gateway` still has `"2593:2593"`
- `yamul-dashboard` has `"8080:8080"`
- `yamul-envoy` has no ports section

- [ ] **Step 3: Commit**

```bash
git add docker-compose.yml
git commit -m "feat: add dashboard and envoy services, lock backend ports to internal"
```

---

## Task 6: Update Flutter gRPC endpoint

**Files:**
- Modify: `yamul-dashboard/lib/core/grpc/login_grpc_service.dart:7-9`

The hardcoded `http://localhost:8091` breaks when running in Docker — the browser on the host cannot reach the container-internal Envoy address. `Uri.base` resolves to the browser's current page URL, so the gRPC endpoint automatically matches whatever host:port the user opened the dashboard on.

- [ ] **Step 1: Edit login_grpc_service.dart**

In `yamul-dashboard/lib/core/grpc/login_grpc_service.dart`, replace:
```dart
    final channel = GrpcWebClientChannel.xhr(
      Uri.parse('http://localhost:8091'),
    );
```
with:
```dart
    final grpcUri = Uri(
      scheme: Uri.base.scheme,
      host: Uri.base.host,
      port: Uri.base.port,
      path: '/grpc/',
    );
    final channel = GrpcWebClientChannel.xhr(grpcUri);
```

- [ ] **Step 2: Verify the file compiles locally (optional but recommended)**

If Flutter is installed locally:
```bash
cd yamul-dashboard
flutter analyze lib/core/grpc/login_grpc_service.dart
```
Expected: no errors.

- [ ] **Step 3: Commit**

```bash
git add yamul-dashboard/lib/core/grpc/login_grpc_service.dart
git commit -m "feat: derive gRPC endpoint from page origin for Docker compatibility"
```

---

## Task 7: Integration smoke test

**Files:** none (verification only)

Verify the full stack starts with `docker-compose up` and only the intended ports are reachable.

- [ ] **Step 1: Start all services**

```bash
docker-compose up --build
```
Expected startup order: `yamul-backend` starts → passes healthcheck (may take 60–100 s for JVM startup) → `yamul-envoy` starts → `yamul-dashboard` starts. `yamul-gateway` starts in parallel.

Watch for:
- `yamul-backend` transitions from `starting` to `healthy` (not `unhealthy`)
- `yamul-envoy` and `yamul-dashboard` containers reach `Up` state

If `yamul-backend` stays `unhealthy` after 120 s, check its logs: `docker-compose logs yamul-backend`

- [ ] **Step 2: Verify dashboard is reachable**

In a separate terminal:
```bash
curl -s -o /dev/null -w "%{http_code}" http://localhost:8080
```
Expected: `200`

- [ ] **Step 3: Verify backend ports are not reachable from host**

```bash
# Each of these should fail (connection refused = port not mapped to host)
curl --connect-timeout 2 -s http://localhost:8087 && echo "FAIL: 8087 exposed" || echo "8087 internal OK"
curl --connect-timeout 2 -s http://localhost:8088 && echo "FAIL: 8088 exposed" || echo "8088 internal OK"
curl --connect-timeout 2 -s http://localhost:8089 && echo "FAIL: 8089 exposed" || echo "8089 internal OK"
```
Expected: all three print `internal OK`.

- [ ] **Step 4: Verify dashboard loads in browser**

Open http://localhost:8080 in a browser.
Expected: Flutter app renders. The browser DevTools console should not show CORS errors or failed gRPC connections on initial load.

- [ ] **Step 5: Stop and clean up**

```bash
docker-compose down
```
