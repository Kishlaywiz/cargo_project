-- +goose Up
-- +goose StatementBegin
CREATE TYPE quote_status AS ENUM(
    'Buy Rate Pending',
    'Sell Rate Pending',
    'Approval Pending',
    'Approved',
    'Rejected',
    'Expired'
);
CREATE table if not exists quotes(
    id uuid PRIMARY Key not null,
    currency text,
    partner text,
    booking_id uuid,
    validity text,
    liner Text,
    transit_days int,
    free_days int,
    origin_date text,
    destination_date text,
    charges_info jsonb,
    remarks text,
    quote_status quote_status,
    created_at timestamp,
    created_by uuid,
    updated_at timestamp,
    updated_by uuid,
    CONSTRAINT booking_id FOREIGN KEY(booking_id) REFERENCES booking(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TYPE quote_status;
DROP TABLE IF EXISTS quotes;
-- +goose StatementEnd