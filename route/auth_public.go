package route

import (
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"git.bumpsoo.dev/bumpsoo/book-api/password"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SignUp(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if len(admin.Username) > 20 || len(admin.Username) < 5 || len(admin.Password) > 20 || len(admin.Password) < 5 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"bad input": "username or password is invalid",
		})
	}
	db, err := database.Connect()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "can not connect to database. " + err.Error(),
		})
	}
	sql := db.QueryRow("SELECT uuid FROM admin WHERE username=$1", admin.Username)
	var row string
	if err := sql.Scan(&row); err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"bad input": "username exists",
		})
	}
	admin.Password = password.HashPassword(admin.Password)
	admin.Id, err = uuid.NewUUID()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	db.QueryRow("INSERT INTO admin VALUES ($1, $2, $3)", admin.Id, admin.Username, admin.Password)
	c.JSON(http.StatusOK, gin.H{
		"sign-up": "success",
		"admin": map[string]any{
			"uuid":     admin.Id.String(),
			"username": admin.Username,
		},
		"links": map[string]string{
			"GET books":    "/v1/book",
			"POST sign-in": "/v1/auth/sign-in",
		},
	})
}

func SignIn(c *gin.Context) {
	var admin model.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	if len(admin.Username) > 20 || len(admin.Username) < 5 || len(admin.Password) > 20 || len(admin.Password) < 5 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"bad input": "username or password is invalid",
		})
	}
	db, err := database.Connect()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "can not connect to database. " + err.Error(),
		})
	}
	row := db.QueryRow("SELECT * FROM admin WHERE username=$1", admin.Username)
	var dbAdmin model.Admin
	if err := row.Scan(&dbAdmin); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"bad input": "username does not exist",
		})
	}
	if dbAdmin.Password != password.HashPassword(admin.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"bad input": "invalid input",
		})
	}
	// return jwt
	c.JSON(http.StatusOK, gin.H{
		"sign-in": "success",
		// "admin": map[string]any{
		// 	"uuid":     admin.Id.String(),
		// 	"username": admin.Username,
		// },
		// "links": map[string]string{
		// 	"GET books":    "/v1/book",
		// 	"POST sign-in": "/v1/auth/sign-in",
		// },
	})
}
