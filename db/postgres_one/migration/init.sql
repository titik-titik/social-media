drop table if exists "user" cascade;
create table "user"
(
    id          uuid primary key,
    name        text,
    username    text unique,
    email       text unique,
    password    text,
    avatar_url  text,
    bio         text,
    is_verified boolean,
    created_at  timestamptz,
    updated_at  timestamptz,
    deleted_at  timestamptz
);

drop table if exists "post" cascade;
create table "post"
(
    id          uuid primary key,
    user_id     uuid,
    image_url   text,
    description text,
    created_at  timestamptz,
    updated_at  timestamptz,
    deleted_at  timestamptz,
    constraint fk_post_user_id_user_id foreign key (user_id) references "user" (id)
);


select *
from "user"
where id = '3a19a20c-ff4f-4f2f-8768-86925bca0339';
