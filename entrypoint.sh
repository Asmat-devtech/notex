#!/bin/sh
echo "=== notex entrypoint starting ==="
echo "PORT=${PORT}"
echo "OPENAI_API_KEY length=$(echo -n "${OPENAI_API_KEY}" | wc -c)"
ls -la ./notex 2>&1 || echo "ERROR: notex binary not found"
export SERVER_PORT="${PORT:-8080}"
echo "Starting notex on SERVER_PORT=${SERVER_PORT}"
exec ./notex -server
