package handler

import (
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	db := database.GetPq()
	rows, err := db.Query("SELECT id, title FROM book")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"book": "no books in table " + err.Error(),
		})
	}
	defer rows.Close()
	var ret map[string]string
	for rows.Next() {
		var i, t string
		err := rows.Scan(&i, &t)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		ret[t] = "/v1/book/" + i
	}
	c.JSON(http.StatusOK, gin.H{
		"book":  "success",
		"links": ret,
	})
}

func GetBook(c *gin.Context) {
	db := database.GetPq()
	var id string
	if err := c.ShouldBindUri(&id); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "id must be passed",
		})
	}
	var book model.Book
	row := db.QueryRow("SELECT * FROM book where id=$1", id)
	if err := row.Scan(&book.Id, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Author); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"book": "no book is selected",
		})
	}
	ret := "v1/book/" + id
	c.JSON(http.StatusOK, gin.H{
		"book": map[string]any{
			"id":         id,
			"created_at": book.CreatedAt,
			"updated_at": book.UpdatedAt,
			"title":      book.Title,
			"author":     book.Author,
		},
		"link": map[string]string{
			"PATCH":  ret,
			"DELETE": ret,
		},
	})
}
