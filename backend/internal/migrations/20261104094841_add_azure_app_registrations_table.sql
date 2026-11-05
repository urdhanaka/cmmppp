-- +goose Up
CREATE TABLE IF NOT EXISTS azure_app_registrations (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  azure_cloud_id UUID NOT NULL,
  client_id UUID NOT NULL,
  client_secret_ref TEXT NOT NULL,
  secret_expires_at TIMESTAMPTZ,

  created_by TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ,
  updated_by TEXT,
  deleted_at TIMESTAMPTZ,

  CONSTRAINT unique_client_id UNIQUE(client_id),
  CONSTRAINT fk_azure_clouds
    FOREIGN KEY (azure_cloud_id)
    REFERENCES azure_clouds(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS azure_app_registrations;
