create table users (
    id serial constraint users_pk primary key,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at timestamp,
    first_name  varchar(255) not null,
    last_name   varchar(255) not null,
    middle_name varchar(255),
    phone       varchar(255) not null
);

create unique index users_id_uindex on users (id);

create unique index users_phone_uindex on users (phone);