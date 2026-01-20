#!/bin/sh

# managed-by: taco

GIT_DIR="$(git rev-parse --git-dir)"
LOCK_FILE="$GIT_DIR/taco.push.lock"

if [ -f "$LOCK_FILE" ]; then
  echo "Push blocked: repo is locked by taco. Run: taco unlock"
  exit 1
fi

exit 0
