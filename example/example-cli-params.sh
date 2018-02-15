#!/usr/bin/env bash

cd $(dirname $0)

../bin/linux-amd64/sql-cli --driver="mysql" \
    --dsn="root:password@/mysql" \
    --filepath="stmt.sql" \
    --statement="create_user_local,grant_all_privs,flush_privs"