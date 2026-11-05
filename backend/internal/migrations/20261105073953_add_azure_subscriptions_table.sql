-- +goose Up
CREATE TABLE IF NOT EXISTS azure_subscriptions (
  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ,
  updated_by TEXT,
  deleted_at TIMESTAMPTZ
);

-- +goose Down
DROP TABLE IF EXISTS azure_subscriptions;
