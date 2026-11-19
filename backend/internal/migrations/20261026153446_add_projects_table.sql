-- +goose Up
CREATE TABLE IF NOT EXISTS projects (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  
  -- foreign key to customer
  customer_id UUID NOT NULL,

  project_no TEXT NOT NULL,
  project_name TEXT NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by TEXT,
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,

  -- constraints
  CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers(id),
  CONSTRAINT projects_id_customer_uq UNIQUE (id, customer_id)
);

-- +goose Down
DROP TABLE IF EXISTS projects;
