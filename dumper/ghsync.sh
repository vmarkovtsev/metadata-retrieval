#!/bin/bash -e

if [ -z "$ORGS" ]; then
  ORGS=athenianco
fi

POSTGRES_PASSWORD=postgres docker-entrypoint.sh postgres &
while ! : &>/dev/null </dev/tcp/127.0.0.1/5432; do
  sleep 1
done
metadata-retrieval ghsync --version 1 --orgs=$ORGS --no-forks --db=postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable
pg_dump -U postgres -d postgres -f /io/$ORGS.sql
killall postgres

