create table if not exists users (
    id serial primary key,
    email text not null,
    password text not null,
    is_admin boolean not null,
    birth_date text not null,
    username text not null
);

create table if not exists message(
    id serial primary key,
    from_id integer references users(id) on delete cascade,
    to_id integer references users(id) on delete cascade,
    text text not null,
    date_time timestamp not null
);

create table if not exists pin(
    id serial primary key,
    description text not null,
    likes_count integer not null,
    dislikes_count integer not null,
    author_id integer references users(id) on delete cascade,
    pin_link text not null
);

create table if not exists saved_pins(
    pin_id integer references pin(id) on delete cascade,
    user_id integer references users(id) on delete cascade
);

create table if not exists comment(
    id serial primary key,
    is_deleted boolean not null,
    pin_id integer references pin(id) on delete cascade,
    text text not null,
    author_id integer references users(id) on delete cascade,
    date_time timestamp not null
);
