package controllers

import (
	"context"
	"github.com/feriyusuf/go-sign/app/forms"
	"github.com/feriyusuf/go-sign/app/helpers"
	"github.com/feriyusuf/go-sign/app/models_mongo"
	"github.com/feriyusuf/go-sign/app/models_pg"
	"github.com/feriyusuf/go-sign/app/models_pg/commons"
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
	var user commons.ComUser
	if err := models_pg.PGDB.Where(" username = ?", bodyJson.Username).First(&user).Error; err == nil {
		c.JSON(403, gin.H{"message": "Username already exist"})
		return
	}

	// Find public role
	var role commons.ComRole
	if err := models_pg.PGDB.Where(" name = ?", models_pg.PublicUser).First(&role).Error; err != nil {
		log.Fatalf("Error get public role %v", err)
		c.JSON(501, gin.H{"message": "Internal Server Error"})
	}

	// Save to SQL database
	user = commons.ComUser{
		Username: bodyJson.Username,
		Name:     bodyJson.Name,
		Password: helpers.GenerateHashPassword([]byte(bodyJson.Password)),
	}
	models_pg.PGDB.Create(&user)

	// Assign as public user
	roleUser := commons.ComRoleUser{
		UserId:   user.ID,
		RoleId:   role.ID,
		IsActive: true,
	}

	result := models_pg.PGDB.Create(&roleUser)
	if result.Error != nil {
		log.Printf("Error %v", result.Error)
	}

	// Positive Response
	c.JSON(201, gin.H{"message": "Success Create User"})
}

func (h *AuthController) Login(c *gin.Context) {
	var bodyJson forms.Login

	if c.BindJSON(&bodyJson) != nil {
		c.JSON(401, gin.H{"message": "username and password are required"})
		c.Abort()
		return
	}

	// Search existing user
	var user commons.ComUser
	if err := models_pg.PGDB.Where(" username = ?", bodyJson.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"message": "ComUser not found"})
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
		log.Fatalf("Error generate token%v", err)
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	// Destroy session (if any)
	err = models_mongo.DestroySession(bodyJson.Username)
	if err != nil {
		log.Printf("Error Destroy Session %s", err)
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	// Save session to mongodb
	err = models_mongo.CreateSession(bodyJson.Username, expiredTime, jwtToken)

	if err != nil {
		log.Printf("Error Destroy Create %s", err)
		c.JSON(501, gin.H{"message": "[Mongo] Something went wrong, please try again later!"})
		return
	}

	c.JSON(200, gin.H{"message": "Login Success!", "token": jwtToken})
}

func (h *AuthController) Logout(c *gin.Context) {
	username := c.GetString("username")

	helpers.CallUserContext(context.Background())

	err := models_mongo.DestroySession(username)
	if err != nil {
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	c.JSON(201, gin.H{"message": "Success logout"})
}

func (h *AuthController) Refresh(c *gin.Context) {
	username := c.GetString("username")
	err := models_mongo.DestroySession(username)

	if err != nil {
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	// Generate token
	jwtToken, expiredTime, err := helpers.GenerateToken(username)

	// Save session to mongodb
	err = models_mongo.CreateSession(username, expiredTime, jwtToken)

	if err != nil {
		c.JSON(501, gin.H{"message": "Something went wrong, please try again later!"})
		return
	}

	c.JSON(200, gin.H{"message": "Success refresh token", "token": jwtToken})
}
