-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  role_name TEXT NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ,
  updated_by TEXT,
  deleted_at TIMESTAMPTZ
);

INSERT INTO roles (role_name, created_by)
VALUES ('Cloud Engineer', 'SYSTEM'), ('DevOps', 'SYSTEM'), ('FinOps', 'SYSTEM');

-- +goose Down
DROP TABLE IF EXISTS roles;
