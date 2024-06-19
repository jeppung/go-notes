package main

import (
	"fmt"

	"github.com/jeppung/go-notes.git/database"
	"github.com/jeppung/go-notes.git/models"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.ConnectToDB()
}
func main() {
	err := database.Client.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		fmt.Println("Migration failed:", err.Error())
	} else {
		fmt.Println("Models migrated successfully")
	}
}
