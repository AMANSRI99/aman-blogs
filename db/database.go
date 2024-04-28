package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

func init() {
	dsn := fmt.Sprintf("user=postgres.gwggjucegqwotijdxfop password=AAWuHbXs2qCItOAC host=aws-0-ap-south-1.pooler.supabase.com port=5432 dbname=postgres")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	database = db
	// Perform a simple query to validate the connection
	if err := db.Raw("SELECT 1").Scan(&struct{}{}).Error; err != nil {
		panic(fmt.Errorf("failed to ping database: %w", err))
	}

	log.Println("Database connection established successfully")
}

func GetDatabase() *gorm.DB {
	return database
}
