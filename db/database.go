package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

// init() functions are special - they are automatically called when the package is initialized, before the main() function runs.
func init() {
	dsn := "user=postgres.xqenmhuzhqykuhugwchp password=<password> host=aws-0-ap-south-1.pooler.supabase.com port=6543 dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db
}

func GetDatabase() *gorm.DB {
	return database
}
