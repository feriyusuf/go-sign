package controllers

import "github.com/gin-gonic/gin"

type AuthController struct{}

func (h *AuthController) Register(c *gin.Context) {
	c.JSON(200, gin.H{"message": "You will be signed up"})
}

func (h *AuthController) Login(c *gin.Context) {
	c.JSON(200, gin.H{"message": "You will be signed in"})
}
