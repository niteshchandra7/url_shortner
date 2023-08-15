CREATE TABLE links (
   id SERIAL primary key,
   link TEXT not null,
   shorten_link TEXT not null,
   created_at TIMESTAMP default now()
);