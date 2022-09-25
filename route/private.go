package route

import (
	"github.com/gin-gonic/gin"
)

func PrivateRoutes(a *gin.Engine) {
	router := a.Group("/v1")
	router.POST("/book", PostBook)
	router.PATCH("/book/:uuid", PatchBook)
	router.DELETE("/book/:uuid", DeleteBook)

	router.POST("/auth/sign-out", SignOut)
	router.POST("/auth/refresh", Refresh)
}
