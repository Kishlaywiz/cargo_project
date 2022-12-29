-- +goose Up
-- +goose StatementBegin
CREATE TYPE booking_status as ENUM('Booking Confirmed','In Transit','Booking Completed','Booking Created');
CREATE table if not exists booking(
    id uuid PRIMARY key not null,
    booked_on text,
    customer_id uuid,
    terms_of_shipment  text,
    incoterms text,
    origin_port text,
    origin_address text,
    destination_port text,
    destination_address text,
    door_pickup boolean,
    door_delivery boolean,
    origin_customs boolean,
    destination_customs boolean,
    cargo_ready_date text,
    cargo_is_dangerous boolean,
    cargo_is_stackable boolean,
    cargo_dimension_unit text,
    cargo_count int,
    cargo_weight  float,
    cargo_length float,
    cargo_height float,
    cargo_width float,
    cargo_hs_code int,
    remarks text,
    booking_status booking_status,
    confirmed_quote uuid,
    CONSTRAINT fk_customer_id FOREIGN KEY(customer_id) REFERENCES config(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE booking_status;
DROP TABLE IF EXISTS booking;
-- +goose StatementEnd

