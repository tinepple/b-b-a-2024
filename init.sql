create table houses (
    id serial primary key,
    address text not null,
    year integer not null,
    developer text,
    created_at timestamp without time zone default NOW(),
    updated_at timestamp without time zone default NOW()
);

create type user_role as enum ('client', 'moderator');

create table users (
    id uuid primary key,
    email text unique not null,
    password text not null,
    role user_role not null
);

create type flat_status as enum ('created', 'approved', 'declined', 'on moderation');

create table flats (
    id serial primary key,
    house_id integer references houses(id) not null,
    status flat_status not null default 'created',
    number integer,
    price integer not null,
    rooms_count integer not null,
    moderator_id uuid references users(id)
);

create index concurrently flats_house_id_index on flats using btree (house_id);

create table house_user_subscriptions (
    house_id int references houses (id) on update cascade on delete cascade,
    user_id uuid references users (id) on update cascade on delete cascade,
    constraint house_user_subscriptions_pkey primary key (house_id, user_id)
)