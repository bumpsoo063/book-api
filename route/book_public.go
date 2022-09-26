package route

import (
	"log"
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	db, err := database.ConnectPq()
	if err != nil {
		log.Fatal(err)
	}
	row := db.QueryRow("SELECT * FROM books")
	var tmp string
	if err := row.Scan(&tmp); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{})
	db.Close()
}

func GetBook(c *gin.Context) {
	db, err := database.ConnectPq()
	if err != nil {
		log.Fatal(err)
	}
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
	db.Close()
}
