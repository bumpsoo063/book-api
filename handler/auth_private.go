package handler

import (
	"context"
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/token"
	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	tok, err := token.Parse(c.Request.Header.Get("authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"sign-out": "unauthroized",
		})
	}
	uuid, err := token.ExtractAccess(tok)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"sign-out": "unauthroized",
		})
	}
	rdb := database.GetRedis()
	_, err = rdb.Del(context.Background(), uuid.String()).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"sign-out": "unauthroized",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"sign-out":       "success",
		"links": map[string]string{
			"GET books": "/v1/book",
		},
	})
}

func Refresh(c *gin.Context) {

}
