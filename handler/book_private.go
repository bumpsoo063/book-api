package handler

import (
	"net/http"
	"time"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostBook(c *gin.Context) {
	book := model.Book{}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"post book": err.Error()})
	}
	if len(book.Author) <= 0 || len(book.Title) <= 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"post book": "bad request"})
	}
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	id, err := uuid.NewUUID()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"post book": "failed to create uuid " + err.Error(),
		})
	}
	book.Id = id.String()
	db := database.GetPq()
	_, err = db.Exec(`INSERT INTO book VALUES ($1, $2, $3, $4, $5)`, book.Id, book.CreatedAt, book.UpdatedAt, book.Title, book.Author)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"post book": "failed to create column " + err.Error()})
	}
	ret := "/v1/book/" + book.Id
	c.JSON(http.StatusOK, gin.H{
		"post book": map[string]any{
			"id":         book.Id,
			"created-at": book.CreatedAt,
			"updated-at": book.UpdatedAt,
			"title":      book.Title,
			"author":     book.Author,
		},
		"link": map[string]string{
			"PATCH":  ret,
			"DELETE": ret,
		},
	})
}

func PatchBook(c *gin.Context) {
	book := model.Book{}
	if err := c.ShouldBindUri(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"patch book": err.Error(),
		})
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"patch book": err.Error(),
		})
	}
	db := database.GetPq()
	var err error
	book.UpdatedAt = time.Now()
	if len(book.Title) > 0 {
		if len(book.Author) > 0 {
			_, err = db.Exec(`UPDATE book SET title = $1, author = $2, updated_at = $3 WHERE id = $4`, book.Title, book.Author, book.UpdatedAt, book.Id)
		} else {
			_, err = db.Exec(`UPDATE book SET title = $1, updated_at = $2 WHERE id = $3`, book.Title, book.UpdatedAt, book.Id)
		}
	} else if len(book.Author) > 0 {
		_, err = db.Exec(`UPDATE book SET author = $1, updated_at = $2 WHERE id = $3`, book.Author, book.UpdatedAt, book.Id)
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"patch book": "title or auther should be input"})
	}
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"patch book": "no column for input id"})
	}
	ret := "/v1/book/" + book.Id
	c.JSON(http.StatusOK, gin.H{
		"patch book": "success",
		"link": map[string]string{
			"GET":    ret,
			"DELETE": ret,
		},
	})
}

func DeleteBook(c *gin.Context) {
	book := model.Book{}
	if err := c.ShouldBindUri(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"delete book": err.Error(),
		})
	}
	db := database.GetPq()
	_, err := db.Exec(`DELETE FROM book WHERE id = $1`, book.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"delete book": "failed to delete column " + err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"delete book": "success",
		"link": map[string]string{
			"GET": "/v1/book",
		},
	})

}
