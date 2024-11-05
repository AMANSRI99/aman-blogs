package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

// init() functions are special - they are automatically called when the package is initialized, before the main() function runs.
func init() {
	dsn := "user=postgres.xqenmhuzhqykuhugwchp password=qf34n6xd@smMF@i host=aws-0-ap-south-1.pooler.supabase.com port=6543 dbname=postgres"

	// Add these parameters to the DSN
	dsn += " sslmode=require application_name=your_app_name"

	config := &gorm.Config{
		PrepareStmt:          false,
		DisableAutomaticPing: true, // Add this
	}

	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		panic(err)
	}

	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	// Clear any existing prepared statements
	if _, err := sqlDB.Exec("DEALLOCATE ALL"); err != nil {
		log.Printf("Warning: Failed to deallocate prepared statements: %v", err)
	}

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
