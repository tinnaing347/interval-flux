#!/bin/bash

set -e

until influx; do
    >&2 echo "Influxdb is unavailable - sleeping"
    sleep 1
done

>&2 echo "Influxdb is up - continuing"

go run main.go

exec "$@"