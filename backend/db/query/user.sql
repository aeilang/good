-- name: GetUsers :many
select * from users
order by id;

-- name: GetUserById :one
select * from users
where id = $1
limit 1;

-- name: GetUserByEmail :one
select * from users
where email = $1
limit 1;

-- name: CreateUser :exec
insert into users (name, email, role, password)
values ($1, $2, $3, $4);


-- name: UpdateRefreshTokenByEmail :exec
update users set
refresh_token = $1
where email = $2;

-- name: UpdatePasswordByEmail :exec
update users set
password = $1
where email = $2;

-- name: DeletUserByEmail :exec
delete from users where email = $1;

