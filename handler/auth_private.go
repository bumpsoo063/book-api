package handler

import (
	"context"
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"git.bumpsoo.dev/bumpsoo/book-api/token"
	"github.com/gin-gonic/gin"
)

func SignOut(c *gin.Context) {
	uuid, err := token.ExtractAccess(c.Request.Header.Get("authorization"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	rdb := database.GetRedis()
	_, err = rdb.Del(context.Background(), uuid).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sign-out": "success",
		"links": map[string]string{
			"GET": "/v1/book",
		},
	})
}

func Refresh(c *gin.Context) {
	var tk model.Token
	if err := c.ShouldBindJSON(&tk); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := token.ExtractRefresh(tk.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	rdb := database.GetRedis()
	userId, err := rdb.Get(context.Background(), id).Result()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	tk, err = token.Generate(userId, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := database.SetToken(userId, &tk, rdb, true); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"refresh":      "success",
		"access_token": tk.AccessToken,
		"links": map[string]string{
			"GET": "/v1/book",
		},
	})
}
