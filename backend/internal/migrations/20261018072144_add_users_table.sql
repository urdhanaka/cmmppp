-- +goose Up
-- create users table
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  user_id TEXT NOT NULL,
  is_platform_admin BOOLEAN DEFAULT FALSE,
  created_by TEXT NOT NULL DEFAULT 'SYSTEM',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  updated_by TEXT,
  deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);

-- +goose Down
-- delete users table
DROP TABLE IF EXISTS users;
