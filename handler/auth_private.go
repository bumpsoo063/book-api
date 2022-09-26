package handler

import (
	"context"
	"fmt"
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/token"
	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	uuid, err := token.ExtractAccess(c.Request.Header.Get("authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"sign-out": "unauthroized",
		})
	}
	rdb := database.GetRedis()
	res, err := rdb.Del(context.Background(), uuid.String()).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"sign-out": "unauthroized",
		})
	}
	fmt.Println(res)
	c.JSON(http.StatusOK, gin.H{
		"sign-out": "success",
		"links": map[string]string{
			"GET books": "/v1/book",
		},
	})
}

func Refresh(c *gin.Context) {

}
