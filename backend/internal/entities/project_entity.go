package entities

import "github.com/google/uuid"

type Project struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	ProjectNo   string    `gorm:"type:string;not null"`
	ProjectName string    `gorm:"type:string;not null"`

	Timestamp
}

func (Project) TableName() string {
	return "projects"
}
