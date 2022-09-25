package main

import (
	"git.bumpsoo.dev/bumpsoo/book-api/route"

	// "git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// db, err := database.Connect()
	// if err != nil {
	// 	println("error")
	// 	log.Panic(err)
	// } else {
	// 	row := db.QueryRow("SELECT * FROM book")
	// 	var str string
	// 	if err := row.Scan(&str); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	println(str)
	// }
	app := gin.Default()
	route.PublicRoutes(app)
	app.Run(":3000")
}
