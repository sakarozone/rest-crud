package usercontrollers

import (
	"learngo/restapiserver/pkg/common/db"
	config "learngo/restapiserver/configs"
	model "learngo/restapiserver/pkg/movies/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	config, err := config.ReadConfig()

	if err != nil {
		log.Fatal("Cannot load config", err.Error())
	}

	//Get email and pass from the body
	var body struct {
		Email    string
		Password string
	}
	err = c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}

	//Check for the user in the database
	var user model.User

	db, err := db.ReturnDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db.First(&user, "email= ?", body.Email)

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	//Check if the password matches

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		return
	}

	//Create a new jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"expires": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenStringVal, err := token.SignedString([]byte(config.SecretKey))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenStringVal,
	})

}
