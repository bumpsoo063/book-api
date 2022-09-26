package route

import (
	"git.bumpsoo.dev/bumpsoo/book-api/handler"
	"git.bumpsoo.dev/bumpsoo/book-api/middleware"
	"github.com/gin-gonic/gin"
)

func PrivateRoutes(a *gin.Engine) {
	router := a.Group("/v1")
	router.POST("/book", handler.PostBook).Use(middleware.Token)
	router.PATCH("/book/:uuid", handler.PatchBook).Use(middleware.Token)
	router.DELETE("/book/:uuid", handler.DeleteBook).Use(middleware.Token)

	router.DELETE("/auth/sign-out", handler.SignOut)
	router.POST("/auth/refresh", handler.Refresh)
}
