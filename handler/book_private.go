package handler

import (
	"fmt"
	"net/http"
	"time"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostBook(c *gin.Context) {
	var book model.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"post book": err.Error(),
		})
	}
	if len(book.Author) <= 0 || len(book.Title) <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"post book": "bad request",
		})
	}
	book.CreatedAt = time.Now().Unix()
	book.UpdatedAt = time.Now().Unix()
	id, err := uuid.NewUUID()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"post book": "failed to create uuid " + err.Error(),
		})
	}
	book.Id = id.String()
	db := database.GetPq()
	res, err := db.Exec(`INSERT INTO book VALUES ($1, $2, $3, $4, $5)`, book.Id, book.CreatedAt, book.UpdatedAt, book.Title, book.Author)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"post book": "failed to create column " + err.Error(),
		})
	}
	fmt.Println(res)
	ret := "v1/book/" + book.Id
	c.JSON(http.StatusOK, gin.H{
		"post book": "success",
		"link": map[string]string{
			"PATCH":  ret,
			"DELETE": ret,
		},
	})
}

func PatchBook(c *gin.Context) {

}

func DeleteBook(c *gin.Context) {

}
