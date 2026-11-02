-- cloud_accounts held exactly one "Azure subscription", "AWS account", or "GCP Projects"
-- +goose Up
CREATE TABLE IF NOT EXISTS cloud_accounts (
  id UUID PRIMARY KEY DEFAULT uuidv7(),

  -- foreign key to customer
  customer_id UUID NOT NULL,

  -- foreign key to cloud_provider
  cloud_provider_id integer NOT NULL,

  display_name TEXT NOT NULL,

  -- constraints
  CONSTRAINT fk_customer FOREIGN KEY (customer_id) REFERENCES customers(id),
  CONSTRAINT fk_cloud_providers FOREIGN KEY (cloud_provider_id) REFERENCES cloud_providers(id)
);

-- +goose Down
DROP TABLE IF EXISTS cloud_accounts;
