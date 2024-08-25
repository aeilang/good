
create table if not exists users
(
    id bigserial primary key,
    created_at timestamp not null default current_timestamp,
    name text not null,
    email citext unique not null,
    password_hash text not null,
    refresh_token text not null default '',
    is_deleted boolean not null,
    version integer not null default 1
);