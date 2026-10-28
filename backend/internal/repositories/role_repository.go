package repositories

import (
	"gorm.io/gorm"
)

type (
	RoleRepository interface{}

	roleRepository struct {
		db *gorm.DB
	}
)

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		db,
	}
}
