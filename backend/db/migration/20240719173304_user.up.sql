create table if not exists "users" (
    id serial primary key,
    name varchar(255) not null,
    email varchar(255) unique not null,
    role varchar(255) not null default 'user',
    password varchar(255) not null,
    refresh_token text,
    created_at timestamp not null default current_timestamp
);