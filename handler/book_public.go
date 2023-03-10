package handler

import (
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/gin-gonic/gin"
)

// @Summary      get books by map{"title": "get link"}
// @Description  all books with link to get details
// @Tags         books
// @Produce      json
// @Success      200  {object} string: string
// @Router       /book [get]
func GetBooks(c *gin.Context) {
	db := database.GetPq()
	rows, err := db.Query(`SELECT id, title FROM book`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	ret := map[string]string{}
	for rows.Next() {
		var i, t string
		err := rows.Scan(&i, &t)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
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
	var book model.Book
	if err := c.ShouldBindUri(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	row := db.QueryRow(`SELECT * FROM book WHERE id = $1`, book.Id)
	if err := row.Scan(&book.Id, &book.CreatedAt, &book.UpdatedAt, &book.Title, &book.Author); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ret := "/v1/book/" + book.Id
	c.JSON(http.StatusOK, gin.H{
		"book": map[string]any{
			"id":         book.Id,
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
