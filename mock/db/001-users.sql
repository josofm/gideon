\c gideondev

CREATE TABLE IF NOT EXISTS "user"
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

INSERT INTO "user" ("name", "sex", "age", "password", "email") VALUES (
		    'gideon jura', 'm','44','$2a$10$mU.3JaxIdQleWHKzubf.yO6n5Ulnizmbju/i73XkuaDTt5lO1fhEC', 'gideon@mtg.com');
