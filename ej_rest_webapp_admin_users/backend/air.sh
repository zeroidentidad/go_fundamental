#!/bin/bash
set -e 

# Always rollback shell options before exiting or returning
trap "set +e" EXIT RETURN

echo "[+] Run container"
docker start goback_mysqldb

echo "[+] Live-reload Air binary"
air -c .air.linux.conf

echo "[+] Stop container on exit..."
docker stop goback_mysqldb