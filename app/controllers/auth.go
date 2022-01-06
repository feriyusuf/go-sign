package controllers

import (
	"github.com/feriyusuf/go-sign/app/forms"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (h *AuthController) Register(c *gin.Context) {
	var data forms.RegisterUserCommand

	if c.BindJSON(&data) != nil {
		c.JSON(406, gin.H{"message": "name, username and password are required"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "You will be signed up"})
}

func (h *AuthController) Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "You will be signed in"})
}
