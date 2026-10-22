package store

import (
	"cmmppp/internal/config"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config config.EnvConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=prefer",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Pass, config.Database.Name,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(15 * time.Minute)
	sqlDB.SetConnMaxLifetime(1 * time.Hour)

	return db, nil
}
