package usercontrollers

import (
	"learngo/restapiserver/pkg/common/db"
	model "learngo/restapiserver/pkg/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//Get email and pass from the body

	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//Create the user
	user := model.User{Email: body.Email, Password: string(hash)}
	db, err := db.ConnectToDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{"error": result.Error.Error()})
		return
	}

	//Send a response
	c.JSON(200, gin.H{"data": "Created successfully"})

}