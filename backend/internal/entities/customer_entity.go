package entities

import "github.com/google/uuid"

type Customer struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuidv7()" json:"id"`
	CustomerName string    `gorm:"type:string;not null" json:"customer_name"`

	Timestamp
}

func (Customer) TableName() string {
	return "customers"
}
