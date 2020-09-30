#!/bin/bash
set -e

until curl -sL -I {$DB_ADDR}ping ; do
    >&2 echo "influx is unavailable - sleeping"
    sleep 1
done

>&2 echo "influx is up"
curl -XPOST {$DB_ADDR}query --data-urlencode 'q=CREATE DATABASE "ivdb"'
go run main.go

exec "$@"