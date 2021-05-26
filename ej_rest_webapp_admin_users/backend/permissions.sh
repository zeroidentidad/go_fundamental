#!/bin/bash
set -e 

# Rollback shell options before exiting or returning
trap "set +e" EXIT RETURN

echo "[+] Copy permissions.sql into container"
docker cp permissions.sql goback_mysqldb:/permissions.sql

echo "[+] Login into container and Setup DB"
docker exec -itd goback_mysqldb /bin/sh -c 'mysql --user=admin --password=passwordx < /permissions.sql 2>/dev/null | grep -v "mysql: [Warning] Using a password on the command line interface can be insecure."'

echo "[+] Success running script :)"