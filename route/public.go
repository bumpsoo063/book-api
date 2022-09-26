package route

import (
	"github.com/gin-gonic/gin"
)

func PublicRoutes(a *gin.Engine) {
	router := a.Group("/v1")
	router.GET("/book", GetBooks)
	router.GET("/book/:uuid", GetBook)
	router.POST("/auth/sign-up", SignUp)
	router.POST("/auth/sign-in", SignIn)
}
