-- +goose Up
-- create users table
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  email TEXT NOT NULL,
  password TEXT NOT NULL,
  name TEXT NOT NULL,
  
  -- user roles (engineer, devops, etc) is defined in another table
  -- because a user can have many different roles
  -- based on the project/customer, but user 
  -- that has access as the platform admin can only
  -- be platform admin, so is_platform_admin column determines
  -- if a user is platform admin and is located on users table
  is_platform_admin BOOLEAN DEFAULT FALSE,

  created_by TEXT NOT NULL DEFAULT 'SYSTEM',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  updated_by TEXT,
  deleted_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_id ON users (id);

-- +goose Down
-- delete users table
DROP TABLE IF EXISTS users;
