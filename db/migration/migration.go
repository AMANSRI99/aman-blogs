package main

import (
	"github.com/AMANSRI99/aman-blogs/db"
	"github.com/AMANSRI99/aman-blogs/pkg/models"
)

func main() {
	db := db.GetDatabase()
	db.Debug()
	db.AutoMigrate(models.Post{})

}
