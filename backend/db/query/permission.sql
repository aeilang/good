-- name: CreatePermission :exec
insert into permissions (code)
values ($1);

-- name: DeletePermissionByCode :exec
delete from permissions
where code = $1;

-- name: UpdatePermissionByCode :exec
update permissions
set code = $1
where code = $2;

-- name: GetPermissionsForUserByEmail :many
select p.code
from permissions p
inner join users_permissions up on up.permission_id = p.id
inner join users u on up.user_id = u.id
where u.email = $1;

-- name: AddPermissionForUserByIdAndCodes :exec
insert into users_permissions (user_id, permission_id)
select @userId::integer, p.id
from permissions p
where p.code = any(@codes::text[]);


-- name: DeletePermissionForUserByIdAndCodes :exec
delete from users_permissions up
where up.user_id = @userId::integer
and up.permission_id = any(
    select p.id from permissions p
    where p.code = @codes::text[]
);


