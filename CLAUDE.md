# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**uo-yamul-server** is a game server implementation for Ultima Online (YAMUL = Yet Another MULtima server). It uses a microservices architecture:

- **Gateway** (`yamul-gateway/`) — Go service on TCP port 2593 that handles the Ultima Online client protocol (encryption, packet parsing) and proxies to backend via gRPC
- **Backend** (`yamul-backend/`) — Kotlin/Spring Boot service hosting 3 gRPC servers (Login :8087, Character :8088, Game :8089)
- **Dashboard** (`yamul-dashboard/`) — Flutter admin UI for game management
- **Proto definitions** (`api-definitions/`) — Shared Protocol Buffer API contracts

## Build & Test Commands

### Gateway (Go)
```bash
cd yamul-gateway
go build -v ./...       # Build
go test -v ./...        # Run all tests
go test -v ./internal/... -run TestName  # Run single test
```

### Backend (Kotlin/Gradle)
```bash
cd yamul-backend
./gradlew build         # Full build (includes googleJavaFormat and detekt)
./gradlew test          # Run tests
./gradlew detekt        # Code quality analysis
./gradlew googleJavaFormat  # Format code
```

### Dashboard (Flutter)
```bash
cd yamul-dashboard
flutter pub get         # Install dependencies
flutter build web       # Build for web
./devBuild.ps1          # Regenerate Dart gRPC stubs (Docker-based protoc)
```

### Docker (all services)
```bash
docker-compose up       # Build and run gateway + backend
```

## Architecture

```
UO Client (TCP :2593)
    │ Multima Protocol (encrypted)
    ▼
Yamul Gateway (Go)
    │ gRPC
    ├── LoginService (Kotlin :8087)
    ├── CharacterService (Kotlin :8088)  ← auth interceptor
    └── GameService (Kotlin :8089)       ← auth interceptor
```

The 3 Kotlin gRPC servers run in the same JVM process, started in `ServiceMain.kt`. Auth is enforced via `BasicAuthInterceptor` on character and game services.

The Gateway's internal flow: TCP connection → `multima/` (decrypt/decompress) → `services/` (gRPC adapters) → backend.

## Regenerating gRPC Code

After editing `.proto` files in `api-definitions/`:

**Gateway (Go)** — from `yamul-gateway/`:
```bash
protoc --go_out=. --go-grpc_out=. ../api-definitions/backend/*.proto
```

**Backend (Kotlin)** — automatic via Gradle protobuf plugin on `./gradlew build`

**Dashboard (Flutter/Dart)** — run `./devBuild.ps1` (uses Docker)

## Key Files

| File | Purpose |
|------|---------|
| `yamul-gateway/main.go` | Gateway entry point, TCP listener |
| `yamul-gateway/internal/transport/multima/` | UO client protocol handling, crypto |
| `yamul-gateway/internal/services/` | gRPC adapter layer to backend |
| `yamul-backend/src/main/kotlin/.../common/controller/ServiceMain.kt` | Backend entry, starts 3 gRPC servers |
| `yamul-backend/src/main/kotlin/.../auth/controller/BasicAuthInterceptor.kt` | gRPC auth middleware |
| `api-definitions/backend/common.proto` | Shared types (Coordinate, Flags, etc.) |
| `yamul-backend/config/detekt.yml` | Kotlin linting rules |

## Backend Code Style

- Uses **Google Java Format** enforced at build time
- **Detekt** for static analysis (config in `yamul-backend/config/detekt.yml`)
- Java 21 (Microsoft JDK 21.0.6 via Gradle toolchain)
- Kotlin 2.1.21 with coroutines for gRPC streaming

## MUL Files

The backend reads Ultima Online's binary MUL data files. The path is configured in `yamul-backend/.env` via the `MULTIMA_MULFILES_DIR` property. MUL parsing code is under `game/controller/domain/mul/`.
