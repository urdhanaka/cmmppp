-- +goose Up
CREATE TABLE IF NOT EXISTS customers (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  customer_name TEXT NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  updated_by TEXT,
  deleted_at TIMESTAMP
);

-- insert dummy customer data
INSERT INTO customers (customer_name) VALUES ('PT. Dummy Bersama');

-- +goose Down
DROP TABLE IF EXISTS customers;
