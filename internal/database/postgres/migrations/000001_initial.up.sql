create table if not exists posts (
    id serial primary key,
    title varchar(30) not null,
    content varchar(255) not null,
    created_at timestamp not null,
    username varchar(30)
)