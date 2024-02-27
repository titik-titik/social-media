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

insert into user (id, name, username, email, password, avatar_url, bio, is_verified, created_at, updated_at, deleted_at)
values ('06d6ec1b-f1ef-46bb-8eca-413e2c5d6d20', 'name0', 'username0', 'email0', 'password0',
        'https://placehold.co/400x400?text=avatar_url0', 'bio0', false, now(), now(), null),
       ('06d6ec1b-f1ef-46bb-8eca-413e2c5d6d21', 'name1', 'username1', 'email1', 'password1',
        'https://placehold.co/400x400?text=avatar_url1', 'bio1', true, now(), now(), null);

select *
from user;
