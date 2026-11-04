package entities

import "github.com/google/uuid"

type CloudProvider struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuidv7()" json:"id"`
	ProviderName string    `gorm:"type:string" json:"provider_name"`
}

func (CloudProvider) TableName() string {
	return "cloud_providers"
}
