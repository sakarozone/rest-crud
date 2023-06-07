package middleware

import (
	"fmt"
	"learngo/restapiserver/pkg/common/db"
	model "learngo/restapiserver/pkg/movies/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	//validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the expiry
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token expired"})
			return
		}

		//Find the user with token
		var user model.User
		fmt.Println("here")
		db, err := db.ConnectToDB()
		if err != nil {
			panic("Failed to connect to database: " + err.Error())
		}

		db.First(&user, "email=?", claims["email"])
		fmt.Println("next here")

		if user.Email == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorised"})
		}

		c.Set("user", user)
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})

	}
	//Continue
	fmt.Println("This message is from the middleware")
	c.Next()
}
