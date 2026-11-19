package entities

import (
	"time"

	"gorm.io/gorm"
)

// Timestamp struct that can be imported by another entity
// DeletedAt is using gorm.DeletedAt so we can utilize gorm's soft delete
// UpdatedAt and UpdatedBy are using pointer data type so gorm knows that it is nullable
type Timestamp struct {
	CreatedBy string         `gorm:"type:string;not null" json:"created_by"`
	CreatedAt time.Time      `gorm:"type:timestamptz;not null" json:"created_at"`
	UpdatedBy *string        `gorm:"type:string" json:"updated_by,omitempty"`
	UpdatedAt *time.Time     `gorm:"type:timestamptz;autoUpdateTime:false" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamptz" json:"deleted_at"`
}
