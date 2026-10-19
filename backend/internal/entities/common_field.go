package entities

import (
	"time"

	"gorm.io/gorm"
)

type Timestamp struct {
	CreatedBy string         `gorm:"type:string;not null;default:SYSTEM" json:"created_by"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:timestamp" json:"updated_at,omitempty"`
	UpdatedBy *string        `gorm:"type:string" json:"updated_by,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp" json:"deleted_at"`
}
