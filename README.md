
# Cross platform command line SQL execution

This project can be used to execute SQL scripts from cli.

Binaries are cross-compiled and checked in for the following platforms:
    
    - windows/amd64
    - linux/amd64
    - darwin/amd64
    
Supported db drivers included in the binaries:

    mysql - for details on DSN see https://github.com/go-sql-driver/mysql/#dsn-data-source-name
    postgres - for details on DSN see https://godoc.org/github.com/lib/pq

# Usage
Given an SQL file containing multiple statements:

```sql
-- name: create_user_local
CREATE USER 'local'@'%' IDENTIFIED BY 'password';

-- name: grant_all_privs
GRANT ALL PRIVILEGES ON *.* TO 'local'@'%' with grant option;

-- name: flush_privs
FLUSH PRIVILEGES;
```
Every statement must have a name tag (`--name:<some name>`).
The name tag will be specified as an argument to the command line tool.

linux env:
```sh
export DSQL_DRIVER="mysql"
export DSQL_DSN="root:password@/mysql"
export DSQL_FILEPATH="stmt.sql"
export DSQL_STATEMENT="create_user_local,grant_all_privs,flush_privs"

sql-cli
```
or
```sh
sql-cli --driver="mysql" /
    --dsn="root:password@/mysql" /
    --filepath="stmt.sql" /
    --statement="create_user_local,grant_all_privs,flush_privs"
```

# Parameters

````commandline
Usage of sql-from-cli:
  -driver string
        Specify the sql driver for the db connection. Currently supported: (mysql | postgres).
  -dsn string
        Specify the data source as string. [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
  -filepath string
        Specify the path to the file that needs to be run on the database.
  -statement string
        Specify the statement tag within the sql file to run only a subset of statements.

````
    
# Build from source

```sh
./fetch-dependencies.sh
./cross_compile_linux_amd64.sh
```

