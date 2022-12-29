-- +goose Up
-- +goose StatementBegin
CREATE TYPE emptype AS ENUM ('Customer','ProcExec','SalesExec','Admin');
CREATE table if not exists config(
    id uuid not null PRIMARY KEY,
    user_name text,
    password  text,
    account_type emptype,
    email	Text,
    mobile	text,
    address	Text,
    country	Text 
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE admin_level1;
Drop TABLE IF EXISTS config;
-- +goose StatementEnd
