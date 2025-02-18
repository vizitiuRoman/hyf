begin;

-- Create user table
create table "user"
(
    id            serial primary key,
    name          varchar(255)        not null,
    last_name     varchar(255)        not null,
    email         varchar(255) unique not null,
    password      varchar(255)        not null,
    created_at    timestamp without time zone not null,
    last_login_at timestamp without time zone
);

commit;