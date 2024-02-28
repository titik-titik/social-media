drop table if exists post cascade;
drop table if exists user cascade;

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

select *
from user;
select *
from post;

delete
from user;
delete
from post;
