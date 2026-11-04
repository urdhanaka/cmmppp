package entities

import (
	"time"

	"github.com/google/uuid"
)

type AzureAppRegistration struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:uuidv7()" json:"id"`
	AzureCloudID    uuid.UUID  `gorm:"column:azure_cloud_id;type:uuid;not null" json:"azure_cloud_id"`
	ClientID        uuid.UUID  `gorm:"column:client_id;type:uuid;not null" json:"client_id"`
	ClientSecretRef string     `gorm:"column:client_secret_ref;type:text;not null" json:"client_secret_ref"`
	SecretExpiresAt *time.Time `gorm:"column:secret_expires_at;type:timestamp" json:"secret_expires_at,omitempty"`

	// AzureCloud here is the association (foreign key)
	// gorm will automatically refer AzureCloudID above to [AzureCloud]
	AzureCloud AzureCloud `gorm:"foreignKey:AzureCloudID;references:ID;" json:"-"`

	Timestamp
}

func (AzureAppRegistration) TableName() string {
	return "azure_app_registrations"
}
