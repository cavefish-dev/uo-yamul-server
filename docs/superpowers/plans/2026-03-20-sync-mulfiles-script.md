# sync_mulfiles.sh Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fill `sync_mulfiles.sh` to rsync UO MUL files from the path configured in `yamul-backend/.env` into the local `mulfiles/` directory used by Docker.

**Architecture:** The script reads `MULTIMA_MULFILES_DIR` from `yamul-backend/.env`, then calls `rsync` to copy that directory's contents into `./mulfiles/`. Docker Compose mounts `./mulfiles` into the backend container at `/mulfiles`, so this script is the bridge between the developer's local UO install and the containerised server.

**Tech Stack:** Bash, rsync, grep/cut for `.env` parsing

---

## File Structure

| File | Action | Responsibility |
|------|--------|---------------|
| `sync_mulfiles.sh` | Modify | Parse `.env`, validate source path, run rsync |

---

### Task 1: Implement sync_mulfiles.sh

**Files:**
- Modify: `sync_mulfiles.sh`

**Context:**
- `yamul-backend/.env` currently contains: `MULTIMA_MULFILES_DIR=c:/repos/uo-server-mul-files`
- `docker-compose.yml` mounts `./mulfiles:/mulfiles` into the backend container
- `./mulfiles/*` is gitignored (binary game data, not committed)
- Script will be run from the repo root via Git Bash or a Unix shell

- [ ] **Step 1: Write the script**

Replace `sync_mulfiles.sh` with:

```bash
#!/usr/bin/env bash
set -euo pipefail

ENV_FILE="yamul-backend/.env"

if [ ! -f "$ENV_FILE" ]; then
  echo "Error: $ENV_FILE not found. Run this script from the repo root." >&2
  exit 1
fi

if ! command -v rsync &>/dev/null; then
  echo "Error: rsync not found. On Windows/MSYS2: pacman -S rsync" >&2
  exit 1
fi

MULFILES_DIR=$(grep -E '^MULTIMA_MULFILES_DIR=' "$ENV_FILE" | cut -d'=' -f2-)

if [ -z "$MULFILES_DIR" ]; then
  echo "Error: MULTIMA_MULFILES_DIR not set in $ENV_FILE" >&2
  exit 1
fi

# Convert Windows-style drive path (c:/...) to MSYS2 Unix path (/c/...)
# so rsync does not misinterpret the colon as an SSH host separator.
if command -v cygpath &>/dev/null; then
  MULFILES_DIR=$(cygpath -u "$MULFILES_DIR")
fi

# Strip any trailing slash before appending one, to avoid double-slash.
MULFILES_DIR="${MULFILES_DIR%/}"

if [ ! -d "$MULFILES_DIR" ]; then
  echo "Error: source directory '$MULFILES_DIR' does not exist or is not mounted." >&2
  exit 1
fi

echo "Syncing mulfiles from '$MULFILES_DIR' -> './mulfiles/'"
mkdir -p mulfiles
# --delete removes files in ./mulfiles/ that no longer exist in the source,
# keeping the destination a faithful mirror of the configured UO install.
rsync -av --delete "$MULFILES_DIR/" mulfiles/
echo "Done."
```

- [ ] **Step 2: Make it executable**

```bash
chmod +x sync_mulfiles.sh
```

- [ ] **Step 3: Smoke-test manually**

Run from repo root:
```bash
./sync_mulfiles.sh
```

Expected: files copied from `c:/repos/uo-server-mul-files/` into `./mulfiles/`, rsync progress printed, script exits 0.

Verify destination:
```bash
ls mulfiles/ | head -20
```

Expected: UO MUL files (e.g. `map0.mul`, `statics0.mul`, etc.)

- [ ] **Step 4: Commit**

```bash
git add sync_mulfiles.sh
git commit -m "feat: add sync_mulfiles.sh to rsync MUL files from .env path"
```
