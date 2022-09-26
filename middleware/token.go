package middleware

import (
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/token"
	"github.com/gin-gonic/gin"
)

func Token(c *gin.Context) {
	_, err := token.ExtractAccess(c.Request.Header.Get("authorization"))
	if (err != nil) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"unauthorized": err.Error(),
		})
	} else {
		c.Next()
	}
}
