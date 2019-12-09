package main

import (
	"go-training-restful/config"
	"go-training-restful/models"
	"go-training-restful/routes"
	"os"
)

func main() {
	e := routes.New()
	port := os.Getenv("PORT")
	InitialMigration()

	e.Logger.Fatal(e.Start(":" + port))
}

func InitialMigration() {
	var db = config.DB
	if !db.HasTable(&models.User{}) {
		db.AutoMigrate(&models.User{})
	}
}
