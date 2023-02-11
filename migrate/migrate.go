package main

import (
	"fmt"
	"github.com/nhannguyen1295/golang-gorm-postgres/initializers"
	"github.com/nhannguyen1295/golang-gorm-postgres/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("? Could not migrate", err)
	}
	fmt.Println("? Migration complete")
}
