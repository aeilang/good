-- name: CreateUser :exec
insert into users 
(name, email, password_hash)
values ($1, $2, $3);

-- name: UpdateUserRefreshTokenById :exec
update users set
refresh_token = $1
where id = $2;

-- name: UpdateUserById :exec
update users set
name = $1, email = $2, password_hash = $3, refresh_token=$4, version = version +1
where id = $5 and version = $6 and is_deleted = false;

-- name: DeleteUserByEmail :exec
update users set
is_deleted = true
where id = $1;

-- name: GetUserById :one
select name, email, password_hash, refresh_token
from users
where is_deleted = false
and id = $1;
