-- +goose Up
CREATE TABLE IF NOT EXISTS azure_app_registrations (
  id UUID PRIMARY KEY DEFAULT uuidv7(),
  azure_cloud_id UUID NOT NULL,
  client_id UUID NOT NULL,
  client_secret_ref TEXT NOT NULL,
  secret_expires_at TIMESTAMP,

  created_by TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  updated_by TEXT,
  deleted_at TIMESTAMP,

  CONSTRAINT unique_client_id UNIQUE(client_id),
  CONSTRAINT fk_azure_clouds
    FOREIGN KEY (azure_cloud_id)
    REFERENCES azure_clouds(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
)

-- +goose Down
SELECT 'down SQL query';
