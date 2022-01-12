package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Create db connection
	database, err := gorm.Open("sqlite3", "test.db")

	// Check that db successfully created
	if err != nil {
		panic("Failed to connect to database!")
	}

	// Migrate db schema using AutoMigrate
	database.AutoMigrate(&Book{})

	// Populate DB w/ our db instance
	DB = database
}
