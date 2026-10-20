package entities

import (
	"time"

	"gorm.io/gorm"
)

// Timestamp struct that can be imported by another entity
// DeletedAt is using gorm.DeletedAt so we can utilize gorm's soft delete
// UpdatedAt and UpdatedBy are using pointer data type so gorm knows that it is nullable
type Timestamp struct {
	CreatedBy string         `gorm:"type:string;not null;default:SYSTEM" json:"created_by"`
	CreatedAt time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"type:timestamp" json:"updated_at,omitempty"`
	UpdatedBy *string        `gorm:"type:string" json:"updated_by,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp" json:"deleted_at"`
}
