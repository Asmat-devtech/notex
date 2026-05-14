#!/bin/sh
# Railway injecte PORT dynamiquement. notex utilise SERVER_PORT.
export SERVER_PORT="${PORT:-8080}"
exec ./notex -server
