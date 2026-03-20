# ── Stage 1: Build Flutter web ────────────────────────────────────────────────
FROM ghcr.io/cirruslabs/flutter:3.41.5 AS flutter-builder

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
