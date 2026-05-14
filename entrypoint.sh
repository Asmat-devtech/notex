#!/bin/sh
# Railway injecte PORT dynamiquement — notex lit SERVER_PORT
export SERVER_PORT="${PORT:-8080}"
exec /app/notex -server
