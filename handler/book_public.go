package handler

import (
	"log"
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	db := database.GetPq()
	row := db.QueryRow("SELECT * FROM books")
	var tmp string
	if err := row.Scan(&tmp); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetBook(c *gin.Context) {
	db := database.GetPq()
	id, f := c.GetQuery("id")
	if f == false {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id must be passed",
		})
		return
	}
	row := db.QueryRow("SELECT * FROM book where id=$1", id)
	var tmp string
	if err := row.Scan(&tmp); err != nil {
		log.Fatal(err)
	}
}
