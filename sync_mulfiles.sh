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
