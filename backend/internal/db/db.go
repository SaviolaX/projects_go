package db

import (
	"log"

	"github.com/SaviolaX/blog/internal/config"
	"github.com/SaviolaX/blog/internal/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Connect(cfg config.DBConfig) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(cfg.Path), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("database connected:", cfg.Path)
	return db
}

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Category{},
	)
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("database migrated successfully")
}
