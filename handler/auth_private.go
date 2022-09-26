package handler

import (
	"context"
	"fmt"
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
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
	res, err := rdb.Del(context.Background(), uuid).Result()
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
	var tk model.Token
	if err := c.ShouldBindJSON(&tk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"refresh": "refresh token must be sent",
		})
	}
	id, err := token.ExtractRefresh(tk.RefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"refresh": "unauthroized",
		})
	}
	rdb := database.GetRedis()
	userId, err := rdb.Get(context.Background(), id).Result()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"refresh": "unauthroized",
		})
	}
	tk, err = token.Generate(userId, true)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"refresh": "failed to generate token " + err.Error(),
		})
	}
	if err := database.SetToken(userId, &tk, rdb); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to set token " + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"refresh": "success",
		"links": map[string]string{
			"GET books": "/v1/book",
		},
	})
}
