begin;

-- Create group table
create table "group"
(
    id         serial primary key,
    group_name varchar(255)             not null,
    user_id    int references "user" (id),
    created_at timestamp with time zone not null
);

-- Create todo table
create table todo
(
    id                serial primary key,
    group_id          int references "group" (id),
    user_id           int                      not null,
    title             text                     not null,
    description       text                     not null,
    location          varchar(255),
    urgency_level     varchar(50),
    created_at        timestamp with time zone not null,
    completion_status boolean
);

-- Create rating table
create table rating
(
    id            serial primary key,
    todo_id       int references todo (id),
    rated_user_id int references "user" (id),
    rating_value  int                      not null,
    created_at    timestamp with time zone not null
);

-- Create group_member table
create table group_member
(
    id             serial primary key,
    group_id       int references "group" (id),
    member_user_id int references "user" (id),
    joined_at      timestamp with time zone not null
);

commit;