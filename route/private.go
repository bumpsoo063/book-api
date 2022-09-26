package route

import (
	"git.bumpsoo.dev/bumpsoo/book-api/handler"
	"git.bumpsoo.dev/bumpsoo/book-api/middleware"
	"github.com/gin-gonic/gin"
)

func PrivateRoutes(a *gin.Engine) {
	router := a.Group("/v1")
	router.POST("/book", middleware.Token, handler.PostBook)
	router.PATCH("/book/:id", middleware.Token, handler.PatchBook)
	router.DELETE("/book/:id", middleware.Token, handler.DeleteBook)

	router.DELETE("/auth/sign-out", handler.SignOut)
	router.POST("/auth/refresh", handler.Refresh)
}
