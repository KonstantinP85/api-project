CREATE TABLE users
(
    id          serial          not null unique,
    name        varchar(255)    not null,
    email       varchar(255)    not null unique,
    password    varchar(255)    not null
);

CREATE TABLE api_list
(
    id           serial          not null unique,
    title        varchar(255)    not null,
    description  varchar(1000)   not null
);

CREATE TABLE users_lists
(
    id           serial                                         not null unique,
    user_id      int references users (id) on delete cascade    not null,
    list_id      int references api_list (id) on delete cascade not null
);