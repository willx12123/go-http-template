#!/bin/bash

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
OUT_DIR="$SCRIPT_DIR/build"

CLIENT_PATH="$SCRIPT_DIR/client"
SERVER_PATH="$SCRIPT_DIR/server"

SERVER_RUNTIME_SCRIPTS_PATH="$SERVER_PATH/scripts/runtime"
SERVER_CONFIG_FILE="$SERVER_PATH/.env"
SERVER_CONFIG_FILE_PROD="$SERVER_PATH/.env.production"
SERVER_EXEC_FILE="$SERVER_PATH/server"

if [ -n "$ONLY" ]; then
  IFS=',' read -ra ONLY_ARRAY <<< "$ONLY"
  for i in "${!ONLY_ARRAY[@]}"; do
    ONLY_ARRAY[$i]=$(echo "${ONLY_ARRAY[$i]}" | xargs)
  done
else
  ONLY_ARRAY=("server" "client")
fi

echo "=====> building target=${ONLY_ARRAY[@]}"

containsElement () {
  local e
  for e in "${@:2}"; do
    if [ "$e" == "$1" ]; then
      return 0
    fi
  done
  return 1
}

cd "$SCRIPT_DIR"
rm -rf "$OUT_DIR"
mkdir -p "$OUT_DIR"

buildClient() {
  cd "$CLIENT_PATH"
  pnpm build
  cp -r "$CLIENT_PATH/dist" "$OUT_DIR/frontend"
}

buildServer() {
  cd "$SERVER_PATH"
  GOOS=linux GOARCH=amd64 go build
  chmod +x "$SERVER_EXEC_FILE"
  cp "$SERVER_RUNTIME_SCRIPTS_PATH/"* "$OUT_DIR"
  cp "$SERVER_EXEC_FILE" "$OUT_DIR"
  cp "$SERVER_CONFIG_FILE" "$OUT_DIR"
  cp "$SERVER_CONFIG_FILE_PROD" "$OUT_DIR"
}

if containsElement "client" "${ONLY_ARRAY[@]}"; then
  echo "=====> start to build client"
  buildClient
  echo "=====> build client done"
fi

if containsElement "server" "${ONLY_ARRAY[@]}"; then
  echo "=====> start to build server"
  buildServer
    echo "=====> build server done"
fi

if [ "${#ONLY_ARRAY[@]}" -ge 2 ]; then
  cd "$SCRIPT_DIR"
fi
