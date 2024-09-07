package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
	}
	err = SeedData(db)
	if err != nil {
		fmt.Println("Failed to seed data")
	}
	repo := NewRepo(db)
	controller := NewController(repo)

	r.GET("/list", controller.List)
	r.GET("/get/:id", controller.Get)
	r.Run(":8080")
}
