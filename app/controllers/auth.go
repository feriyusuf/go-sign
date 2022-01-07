package controllers

import (
	"github.com/feriyusuf/go-sign/app/forms"
	"github.com/feriyusuf/go-sign/app/helpers"
	"github.com/feriyusuf/go-sign/app/models_pg"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (h *AuthController) Register(c *gin.Context) {
	var input forms.Register

	// Body json validation
	if c.BindJSON(&input) != nil {
		c.JSON(406, gin.H{"message": "name, username and password are required"})
		c.Abort()
		return
	}

	// Search existing user
	var user models_pg.User
	if err := models_pg.PGDB.Where(" username = ?", input.Username).First(&user).Error; err == nil {
		c.JSON(403, gin.H{"message": "Username already exist"})
		return
	}

	// Save to SQL database
	user = models_pg.User{
		Username: input.Username,
		Name:     input.Name,
		Password: helpers.GenerateHashPassword([]byte(input.Password)),
	}
	models_pg.PGDB.Create(&user)

	// Positive Response
	c.JSON(201, gin.H{"message": "Success registration"})
}

func (h *AuthController) Login(c *gin.Context) {
	var input forms.Login

	if c.BindJSON(&input) != nil {
		c.JSON(401, gin.H{"message": "username and password are required"})
		c.Abort()
		return
	}

	// Search existing user
	var user models_pg.User
	if err := models_pg.PGDB.Where(" username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"message": "User not found"})
		return
	}

	// Compare registered password and login password
	if err := helpers.ComparePassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(401, gin.H{"message": "Incorrect Password"})
		return
	}

	// TODO: GenerateHashPassword token

	c.JSON(200, gin.H{"message": "You will be signed in"})
}
