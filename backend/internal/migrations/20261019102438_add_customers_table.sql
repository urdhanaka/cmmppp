-- +goose Up
CREATE TABLE IF NOT EXISTS customers (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  customer_name TEXT NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by TEXT,
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

-- insert dummy customer data
INSERT INTO customers (customer_name, created_by) VALUES ('PT. Dummy Bersama', 'SYSTEM');

-- +goose Down
DROP TABLE IF EXISTS customers;
