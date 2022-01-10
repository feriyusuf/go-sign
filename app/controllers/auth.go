package controllers

import (
	"github.com/feriyusuf/go-sign/app/forms"
	"github.com/feriyusuf/go-sign/app/helpers"
	"github.com/feriyusuf/go-sign/app/models_mongo"
	"github.com/feriyusuf/go-sign/app/models_pg"
	"github.com/gin-gonic/gin"
	"log"
)

type AuthController struct{}

func (h *AuthController) Register(c *gin.Context) {
	var bodyJson forms.Register

	// Body json validation
	if c.BindJSON(&bodyJson) != nil {
		c.JSON(406, gin.H{"message": "name, username and password are required"})
		c.Abort()
		return
	}

	// Search existing user
	var user models_pg.User
	if err := models_pg.PGDB.Where(" username = ?", bodyJson.Username).First(&user).Error; err == nil {
		c.JSON(403, gin.H{"message": "Username already exist"})
		return
	}

	// Save to SQL database
	user = models_pg.User{
		Username: bodyJson.Username,
		Name:     bodyJson.Name,
		Password: helpers.GenerateHashPassword([]byte(bodyJson.Password)),
	}
	models_pg.PGDB.Create(&user)

	// Positive Response
	c.JSON(201, gin.H{"message": "Success registration"})
}

func (h *AuthController) Login(c *gin.Context) {
	var bodyJson forms.Login

	if c.BindJSON(&bodyJson) != nil {
		c.JSON(401, gin.H{"message": "username and password are required"})
		c.Abort()
		return
	}

	// Search existing user
	var user models_pg.User
	if err := models_pg.PGDB.Where(" username = ?", bodyJson.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"message": "User not found"})
		return
	}

	// Compare registered password and login password
	if err := helpers.ComparePassword([]byte(user.Password), []byte(bodyJson.Password)); err != nil {
		c.JSON(401, gin.H{"message": "Incorrect Password"})
		return
	}

	// Generate token
	jwtToken, expiredTime, err := helpers.GenerateToken(bodyJson.Username)

	if err != nil {
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	// Destroy session (if any)
	err = models_mongo.DestroySession(bodyJson.Username)
	if err != nil {
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	// Save session to mongodb
	err = models_mongo.CreateSession(bodyJson.Username, expiredTime, jwtToken)

	if err != nil {
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	c.JSON(200, gin.H{"message": "Login Success!", "token": jwtToken})
}

func (h *AuthController) Logout(c *gin.Context) {
	headerToken := c.Request.Header.Get("token")

	// There's no headers' token
	if headerToken == "" {
		c.JSON(401, gin.H{"message": "Token is required"})
		return
	}

	username, err := helpers.DecodeToken(headerToken)

	// Unrecognized token
	if err != nil {
		c.JSON(401, gin.H{"message": "Unknown token"})
		return
	}

	isSessionActive, _ := models_mongo.IsActiveSession(headerToken)

	log.Printf("is Active %b", isSessionActive)

	if !isSessionActive {
		c.JSON(401, gin.H{"message": "Unknown token"})
		return
	}

	err = models_mongo.DestroySession(username)
	if err != nil {
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	c.JSON(201, gin.H{"message": "Success logout"})
}
