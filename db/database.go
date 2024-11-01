package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var database *gorm.DB

// init() functions are special - they are automatically called when the package is initialized, before the main() function runs.
func init() {
	dsn := "user=postgres.xqenmhuzhqykuhugwchp password=qf34n6xd@smMF@i host=aws-0-ap-south-1.pooler.supabase.com port=6543 dbname=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Test connection
    sqlDB, err := db.DB()
    if err != nil {
        log.Fatalf("Failed to get database instance: %v", err)
    }

    // Test the connection
    err = sqlDB.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    log.Println("Successfully connected to database")
    database = db
}

func GetDatabase() *gorm.DB {
	return database
}
