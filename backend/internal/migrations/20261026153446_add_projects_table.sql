-- +goose Up
CREATE TABLE IF NOT EXISTS projects (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  project_no TEXT,
  project_name TEXT NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  updated_by TEXT,
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS projects;
