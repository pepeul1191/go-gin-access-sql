package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("db/app.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
