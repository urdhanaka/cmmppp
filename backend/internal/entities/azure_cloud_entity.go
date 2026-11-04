package entities

import "github.com/google/uuid"

type AzureCloud struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuidv7()" json:"id"`
	CustomerID uuid.UUID `gorm:"column:customer_id;type:uuid;not null;index" json:"customer_id"`
	TenantID   uuid.UUID `gorm:"type:uuid" json:"tenant_id"`
	DomainName string    `gorm:"type:text;not null" json:"domain_name"`

	// Customer here is the association (foreign key)
	// gorm will automatically refer CustomerID above to [Customer]
	Customer Customer `gorm:"foreignKey:CustomerID;references:ID;" json:"-"`

	Timestamp
}

func (AzureCloud) TableName() string {
	return "azure_clouds"
}
