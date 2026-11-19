-- +goose Up
CREATE TABLE IF NOT EXISTS project_cloud_accounts (
  id UUID PRIMARY KEY DEFAULT uuidv7(),

  project_id UUID NOT NULL,
  cloud_account_id UUID NOT NULL,
  customer_id UUID NOT NULL,

  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_by TEXT,
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ,

  CONSTRAINT project_cloud_accounts_project_account_uq UNIQUE (project_id, cloud_account_id),
  CONSTRAINT fk_project_customer FOREIGN KEY (project_id, customer_id)
    REFERENCES projects(id, customer_id) ON DELETE CASCADE,
  CONSTRAINT fk_cloud_account_customer FOREIGN KEY (cloud_account_id, customer_id)
    REFERENCES cloud_accounts(id, customer_id) ON DELETE CASCADE,
  CONSTRAINT fk_customers FOREIGN KEY (customer_id) REFERENCES customers(id)
);

CREATE INDEX IF NOT EXISTS project_cloud_accounts_account_project_idx
  ON project_cloud_accounts (cloud_account_id, project_id);

-- +goose Down
DROP TABLE IF EXISTS project_cloud_accounts;
