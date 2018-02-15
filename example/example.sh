#!/usr/bin/env bash

cd $(dirname $0)

export DSQL_DRIVER="mysql"
export DSQL_DSN="root:password@/mysql"
export DSQL_FILEPATH="stmt.sql"
export DSQL_STATEMENT="create_user_local,grant_all_privs,flush_privs"

../bin/linux-amd64/sql-cli