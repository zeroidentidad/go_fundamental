#!/bin/bash
set -e 

# Rollback shell options before exiting or returning
trap "set +e" EXIT RETURN

echo "[+] Run container"
docker start goback_mysqldb

echo "[+] Wait to container healthy status..."
./wait-hc.sh goback_mysqldb

echo "[+] Wait 30s more please to db ready connections..."
sleep 30

echo "[+] Live-reload Air binary"
air -c .air.linux.conf

echo "[+] Stop container on exit..."
docker stop goback_mysqldb