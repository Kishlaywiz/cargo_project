-- +goose Up
-- +goose StatementBegin
CREATE TYPE status as Enum('Pending', 'Completed');
CREATE table IF not exists task(
    id uuid,
    booking_id uuid,
    name text,
    sub_task text,
    info text,
    status status,
    created_at int,
    created_by uuid,
    updated_at int,
    updated_by uuid,
    CONSTRAINT fk_booking_id FOREIGN KEY (booking_id) REFERENCES booking(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop TABLE if EXISTS task;
DROP TYPE status;
-- +goose StatementEnd
