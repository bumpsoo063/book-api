package route

import (
	"github.com/gin-gonic/gin"
)

func PrivateRoutes(a *gin.Engine) {
	router := a.Group("/v1")
	router.POST("/book", PostBook)
	router.PATCH("/book", PatchBook)
	router.DELETE("/book", DeleteBook)

	router.POST("/auth/sign-out", SignOut)
	router.POST("/auth/refresh", Refresh)
}
