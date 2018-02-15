..\bin\windows-amd64\sql-cli.exe^
 -driver="mysql"^
 -dsn="root:password@/mysql"^
 -filepath="stmt.sql"^
 -statement="create_user_local,grant_all_privs,flush_privs"