package handler

import (
	"net/http"
	"time"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func PostBook(c *gin.Context) {
	book := model.Book{}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()
	id, err := uuid.NewUUID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	book.Id = id.String()
	db := database.GetPq()
	_, err = db.Exec(`INSERT INTO book VALUES ($1, $2, $3, $4, $5)`, book.Id, book.CreatedAt, book.UpdatedAt, book.Title, book.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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
