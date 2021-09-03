package database

import (
	"fmt"
	"github.com/hostingvk4/badgerList/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s port=%s", cfg.Host, cfg.Username, cfg.DbName, cfg.SSLMode, cfg.Password, cfg.Port)
	fmt.Println(dbUri)
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.Debug().AutoMigrate(&models.User{}, &models.List{})

	return db, nil
}
