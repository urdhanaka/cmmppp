package entities

import "github.com/google/uuid"

type Role struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuidv7()" json:"id"`
	RoleName string    `gorm:"type:string"`
}

func (Role) TableName() string {
	return "roles"
}
