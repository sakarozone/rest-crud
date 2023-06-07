package usercontrollers

import (
	"learngo/restapiserver/pkg/common/db"
	model "learngo/restapiserver/pkg/movies/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	//Get email and pass from the body
	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)

	//Check for the user in the database
	var user model.User

	db, err := db.ConnectToDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	db.First(&user, "email= ?", body.Email)

	if user.Email == "" {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	//Check if the password matches

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	//Create a new jwt token

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   user.Email,
		"expires": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenStringVal, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token": tokenStringVal,
	})

}
