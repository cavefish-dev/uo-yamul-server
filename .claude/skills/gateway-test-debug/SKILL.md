---
name: gateway-test-debug
description: Trigger whenever the user wants to run, debug, or inspect yamul-gateway Go tests; also trigger when the user mentions test flags, integration tests, GRPC debug logs, or a specific test name in the gateway.
---

## When to use this skill

Use this skill when the user wants to:
- Run any yamul-gateway Go tests
- Debug a failing gateway test
- Run integration tests for the gateway
- Enable GRPC verbose logging during tests
- Run a specific test by name in the gateway

## Step 1 – Verify permission

Check that `.claude/settings.local.json` contains `"Bash(bash *run-tests.sh*)"` in `permissions.allow`. If it is missing, add it before proceeding.

## Step 2 – Run tests

All gateway tests are run via the wrapper script from the repo root:

```bash
bash .claude/skills/gateway-test-debug/scripts/run-tests.sh <flags>
```

The script `cd`s into `yamul-gateway/` and delegates to `go test "$@"`, so any standard `go test` flags work.

## Common flag recipes

| Goal | Command |
|------|---------|
| All tests, verbose | `bash .claude/skills/gateway-test-debug/scripts/run-tests.sh -v ./...` |
| Single test | `bash .claude/skills/gateway-test-debug/scripts/run-tests.sh -v -run TestFunctionName ./...` |
| Integration tests | `bash .claude/skills/gateway-test-debug/scripts/run-tests.sh -tags integration -v -timeout 30s ./integration/...` |
| GRPC debug logs | `GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info bash .claude/skills/gateway-test-debug/scripts/run-tests.sh -tags integration -v ./integration/...` |
| Race detector | `bash .claude/skills/gateway-test-debug/scripts/run-tests.sh -race -v ./...` |

## Step 3 – Interpret failures

- Build errors → fix import or syntax issues before re-running
- `FAIL` with panic → read the stack trace to identify the goroutine and file
- Integration test timeout → check backend gRPC servers are reachable; increase `-timeout` if needed
- GRPC errors → enable GRPC debug env vars (see recipe above) and re-run to capture wire-level logs
