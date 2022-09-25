package route

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"
	"os"

	"git.bumpsoo.dev/bumpsoo/book-api/database"
	"git.bumpsoo.dev/bumpsoo/book-api/model"
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
	salt := []byte(os.Getenv("SHA_SALT"))
	password := []byte(admin.Password)
	hasher := sha512.New()
	password = append(password, salt...)
	hasher.Write(password)
	hashedPasswordBytes := hasher.Sum(nil)
	admin.Password = hex.EncodeToString(hashedPasswordBytes)
	admin.Id, err = uuid.NewUUID()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	db.QueryRow("INSERT INTO admin VALUES ($1, $2, $3)", admin.Id, admin.Username, admin.Password)
	c.JSON(http.StatusOK, gin.H{
		"sign-up": "success",
	})
}

func SignIn(c *gin.Context) {
	// var admin model.Admin
	// if err := c.ShouldBindJSON(&admin); err != nil {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }
	// if len(admin.Username) > 20 || len(admin.Username) < 5 || len(admin.Password) > 20 || len(admin.Password) < 5 {
	// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	// 		"bad input": "username or password is invalid",
	// 	})
	// }
	// db, err := database.Connect()
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"error": "can not connect to database. " + err.Error(),
	// 	})
	// }
}
