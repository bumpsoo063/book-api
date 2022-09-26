package main

import (
	"log"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/route"

	// "git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := database.ConnectPq(); err != nil {
		log.Fatal("Can not connect to database " + err.Error())
	}
	database.ConnectRedis()
	defer database.GetPq().Close()
	defer database.GetRedis().Close()
	app := gin.Default()
	route.PublicRoutes(app)
	route.PrivateRoutes(app)
	app.Run(":3000")
}
