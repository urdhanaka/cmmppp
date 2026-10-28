-- +goose Up
CREATE TABLE IF NOT EXISTS cloud_providers (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  cloud_provider_name TEXT NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  updated_by TEXT,
  deleted_at TIMESTAMP
);

INSERT INTO cloud_providers (cloud_provider_name, created_by)
VALUES ('Azure', 'SYSTEM'), ('GCP', 'SYSTEM'), ('AWS', 'SYSTEM');

-- +goose Down
DROP TABLE IF EXISTS cloud_providers;
