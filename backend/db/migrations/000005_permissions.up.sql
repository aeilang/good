create table if not exists permissions
(
    id bigserial primary key,
    code text not null
);

create table if not exists users_permissions
(
    user_id bigint not null references users (id) on delete cascade,
    permission_id bigint not null references permissions (id) on delete cascade,
    primary key (user_id, permission_id)
);

insert into permissions (code)
values ('jobs:read'),
('movies:write');