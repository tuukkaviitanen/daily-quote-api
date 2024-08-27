package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB

func ConnectToDatabase() error {
	database, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("database initialization failed: %v", err.Error())
	}

	Database = database
	return nil
}
