-- +goose Up
CREATE TABLE IF NOT EXISTS cloud_providers (
  id integer PRIMARY KEY,
  cloud_provider_name TEXT NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by TEXT,
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

INSERT INTO cloud_providers (id, cloud_provider_name, created_by)
VALUES (1, 'Azure', 'SYSTEM'), (2, 'AWS', 'SYSTEM'), (3, 'GCP', 'SYSTEM');

-- +goose Down
DROP TABLE IF EXISTS cloud_providers;
