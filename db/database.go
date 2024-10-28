package db

import (
	"fmt"

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
}

func GetDatabase() *gorm.DB {
	return database
}
