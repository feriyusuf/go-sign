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

	if c.BindJSON(&input) != nil {
		c.JSON(406, gin.H{"message": "name, username and password are required"})
		c.Abort()
		return
	}

	// TODO: Check username exist

	user := models_pg.User{
		Username: input.Username,
		Name:     input.Name,
		Password: helpers.Generate([]byte(input.Password)),
	}

	models_pg.PGDB.Create(&user)

	c.JSON(201, gin.H{"message": "Success registration"})
}

func (h *AuthController) Login(c *gin.Context) {
	var input forms.Login

	if c.BindJSON(&input) != nil {
		c.JSON(401, gin.H{"message": "username and password are required"})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"message": "You will be signed in"})
}
