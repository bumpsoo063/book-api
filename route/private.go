package route

import (
	"git.bumpsoo.dev/bumpsoo/book-api/handler"
	"github.com/gin-gonic/gin"
)

func PrivateRoutes(a *gin.Engine) {
	router := a.Group("/v1")
	router.POST("/book", handler.PostBook)
	router.PATCH("/book/:uuid", handler.PatchBook)
	router.DELETE("/book/:uuid", handler.DeleteBook)

	router.POST("/auth/sign-out", handler.SignOut)
	router.POST("/auth/refresh", handler.Refresh)
}
