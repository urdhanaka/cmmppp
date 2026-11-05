-- +goose Up
-- make sure citext extension is enabled
CREATE EXTENSION IF NOT EXISTS citext;

-- create users table
CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY DEFAULT uuidv7(),

  -- email and password will be used
  -- as the credential
  email CITEXT NOT NULL UNIQUE,
  password TEXT NOT NULL,
  name TEXT NOT NULL,

  -- user_type either contains 'internal' or 'customer'
  user_type TEXT NOT NULL DEFAULT 'internal',
  
  -- user roles (engineer, devops, etc) is defined in another table
  -- because a user can have many different roles
  -- based on the project/customer, but user 
  -- that has access as the platform admin can only
  -- be platform admin, so is_platform_admin column determines
  -- if a user is platform admin and is located on users table
  is_platform_admin BOOLEAN DEFAULT FALSE,

  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ,
  updated_by TEXT,

  -- is user active can be derived from deleted_at column
  -- if a row has deleted_at column not null, then the user is not active
  deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);
CREATE INDEX IF NOT EXISTS idx_users_id ON users (id);

-- +goose Down
-- delete users table
DROP TABLE IF EXISTS users;
