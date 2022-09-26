package route

import (
	"git.bumpsoo.dev/bumpsoo/book-api/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func PublicRoutes(a *gin.Engine) {
	a.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router := a.Group("/v1")
	router.GET("/book", handler.GetBooks)
	router.GET("/book/:id", handler.GetBook)
	router.POST("/auth/sign-up", handler.SignUp)
	router.POST("/auth/sign-in", handler.SignIn)
}
