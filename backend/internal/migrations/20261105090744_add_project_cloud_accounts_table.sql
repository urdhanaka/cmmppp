-- +goose Up
CREATE TABLE IF NOT EXISTS project_cloud_accounts (
  id UUID PRIMARY KEY DEFAULT uuidv7(),

  -- foreign key to projects
  project_id UUID NOT NULL,

  -- foreign key to cloud_accounts
  cloud_account_id UUID NOT NULL,

  -- foreign key to customer
  -- this looks redundant but two separate FKs cannot compare each other
  -- so it is prone to cross-tenant leaks
  customer_id UUID NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ,
  updated_by TEXT,
  deleted_at TIMESTAMPTZ,

  CONSTRAINT fk_projects FOREIGN KEY (project_id) REFERENCES projects(id),
  CONSTRAINT fk_cloud_accounts FOREIGN KEY (cloud_account_id) REFERENCES cloud_accounts(id),
  CONSTRAINT fk_customers FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- +goose Down
DROP TABLE IF EXISTS project_cloud_accounts;
