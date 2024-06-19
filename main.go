package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeppung/go-notes.git/database"
	"github.com/jeppung/go-notes.git/routes"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	database.ConnectToDB()
}

func main() {
	r := gin.Default()

	routes.Auth(r)

	r.Run()
}
