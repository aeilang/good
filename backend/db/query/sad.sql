
-- name: GetSads :many
select * from sads;

-- name: GetSadById :one
select * from sads
where id = $1
limit 1;

-- name: GetSadByUserId :many
select * from sads
where user_id in (
    select u.id from users u
    where u.id = $1
);

-- name: UpdateSad :exec
update sads set 
name = $1,
reason = $2
where id = $3;

-- name: DeleteSadById :exec
delete from sads where id = $1;

-- name: CreateSad :exec
insert into sads (user_id, name, reason) 
values ($1, $2, $3);

