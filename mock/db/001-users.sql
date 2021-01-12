CREATE DATABASE gideondev;

\c gideondev

CREATE TABLE IF NOT EXISTS user
(
    id       serial  not null
        constraint "user_pkey"
            primary key,
    name     varchar not null,
    sex      varchar not null,
    age      varchar not null,
    password varchar not null,
    email    varchar
);

INSERT INTO user ("id", "name", "sex", "age", "password", "email") VALUES (
		    '1', 'gideon jura', 'm','44','$2a$10$TLRvY9nBsji1snSGJmvBgOETCz8s37hJRPVtvcX4A3iU0XF3eViVq', 'gideon@mtg.com');
