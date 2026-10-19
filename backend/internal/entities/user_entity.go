package entities

import "github.com/google/uuid"

type User struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Email           string    `gorm:"type:string;not null" json:"email"`
	Password        string    `gorm:"type:string;not null" json:"password"`
	UserID          string    `gorm:"type:string;not null" json:"user_id"`
	IsPlatformAdmin bool      `gorm:"type:bool;not null;default:false"`

	Timestamp
}

func (User) TableName() string {
	return "users"
}
