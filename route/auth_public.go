package route

import (
	"net/http"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
	"git.bumpsoo.dev/bumpsoo/book-api/password"
	"git.bumpsoo.dev/bumpsoo/book-api/token"
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
	db, err := database.ConnectPq()
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
	if _, err := db.Exec(`INSERT INTO admin VALUES ($1, $2, $3)`, admin.Id, admin.Username, admin.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"sign-up": "success",
		"admin": map[string]any{
			"uuid":     admin.Id.String(),
			"username": admin.Username,
		},
		"links": map[string]string{
			"GET books":    "/v1/book",
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
	db, err := database.ConnectPq()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "can not connect to database. " + err.Error(),
		})
	}
	row := db.QueryRow(`SELECT * FROM admin WHERE username=$1`, admin.Username)
	var dbAdmin model.Admin
	if err := row.Scan(&dbAdmin.Id, &dbAdmin.Username, &dbAdmin.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"bad input": "username does not exist",
		})
	}
	if dbAdmin.Password != password.HashPassword(admin.Password) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"bad input": "password does not match",
		})
	}
	token, err := token.GenerateToken(dbAdmin.Id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token " + err.Error(),
		})
	}
	rdb := database.ConnectRedis()
	if err := database.SetToken(dbAdmin.Id, &token, rdb); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "failed to set token " + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"sign-in":       "success",
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"admin": map[string]any{
			"uuid":     dbAdmin.Id.String(),
			"username": dbAdmin.Username,
		},
		"links": map[string]string{
			"GET books":     "/v1/book",
		},
	})
}
