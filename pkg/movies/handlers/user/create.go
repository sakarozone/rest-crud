package usercontrollers

import (
	"learngo/restapiserver/pkg/common/db"
	model "learngo/restapiserver/pkg/movies/models"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	//Get email and pass from the body
	var body struct {
		Email    string
		Password string
	}
	err := c.Bind(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Payload"})
		return
	}

	//Check if email and password are valid
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(body.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	if len(body.Password) < 8 || !containsSpecialChar(body.Password) || !containsUpperCase(body.Password) || !containsNumber(body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must have length of 8 characters, one uppercase, one number and one special charater"})
		return
	}

	//Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Create the user
	user := model.User{Email: body.Email, Password: string(hash)}
	db, err := db.ReturnDB()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	result := db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	//Send a response
	c.JSON(http.StatusOK, gin.H{"data": "Created successfully"})

}

// {
//     "email":"sakar@test.com",
//     "password":"testA@LK?lasdf11k?>?@"
// }
