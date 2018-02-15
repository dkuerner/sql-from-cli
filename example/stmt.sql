-- name: create_user_local
CREATE USER 'local'@'%' IDENTIFIED BY 'password';

-- name: grant_all_privs
GRANT ALL PRIVILEGES ON *.* TO 'local'@'%' with grant option;

-- name: flush_privs
FLUSH PRIVILEGES;