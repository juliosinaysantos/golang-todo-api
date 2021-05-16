-- migrate:up
create table if not exists users (
    id serial,
    username character varying (40) unique,
    password character varying (100) not null,
    email character varying (100) unique,
    primary key (id)
);

-- migrate:down
drop table if exists users;
