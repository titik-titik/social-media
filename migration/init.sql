drop table if exists user;
create table user
(
    id          uuid primary key,
    name        text,
    username    text unique,
    email       text unique,
    password    text,
    avatar_url  text,
    bio         text,
    is_verified boolean,
    created_at  datetime,
    updated_at  datetime,
    deleted_at  datetime
);

drop table if exists post;
create table post
(
    id          uuid primary key,
    user_id     uuid references user (id),
    image_url   text,
    description text,
    created_at  datetime,
    updated_at  datetime,
    deleted_at  datetime
);
