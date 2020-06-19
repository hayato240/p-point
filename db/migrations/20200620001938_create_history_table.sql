-- +goose Up
create table point_history
(
    id serial PRIMARY KEY,
    user_id int NOT NULL,
    amount int NOT NULL DEFAULT 0,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose Down
drop table point_history;