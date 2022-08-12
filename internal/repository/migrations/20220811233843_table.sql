-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password varchar(255) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
