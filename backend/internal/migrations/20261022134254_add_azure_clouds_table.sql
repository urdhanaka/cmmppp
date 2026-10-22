-- +goose Up
CREATE TABLE IF NOT EXISTS azure_clouds (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  customer_id UUID NOT NULL,
  tenant_id UUID NOT NULL,
  domain_name TEXT NOT NULL,
  created_by TEXT NOT NULL DEFAULT 'SYSTEM',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  updated_by TEXT,
  deleted_at TIMESTAMP,

  -- constraints
  CONSTRAINT unique_tenant_id UNIQUE(tenant_id),
  CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers(id)
);

-- +goose statementbegin
DO $$
DECLARE
    own_customer_id UUID;
BEGIN
  SELECT id
  INTO own_customer_id
  FROM customers
  LIMIT 1;

  INSERT INTO azure_clouds (customer_id, tenant_id, domain_name)
  VALUES (own_customer_id, 'b6ebaf19-1aa0-434a-81ac-f887f4d4614c', 'dummy.onmicrosoft.com');
END $$;
-- +goose statementend

-- +goose Down
DROP TABLE IF EXISTS azure_clouds;
