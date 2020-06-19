-- +goose Up
create table users
(
    id serial PRIMARY KEY,
    amount int NOT NULL DEFAULT 0,
    date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
drop table users;