package middleware

import (
	"fmt"
	"learngo/restapiserver/pkg/common/db"
	config "learngo/restapiserver/pkg/movies/configs"
	model "learngo/restapiserver/pkg/movies/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
		return
	}
	config, err := config.ReadConfig()

	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}
	//validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.SecretKey), nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("HERE")
		//check the expiry
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token expired"})
			return
		}

		//Find the user with token
		var user model.User
		fmt.Println("here")
		db, err := db.ReturnDB()
		if err != nil {
			panic("Failed to connect to database: " + err.Error())
		}

		db.First(&user, "email=?", claims["email"])
		fmt.Println("next here")

		if user.Email == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
		}

		c.Set("user", user)
	} else {
		fmt.Println("HERE1")
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.Next()
}
