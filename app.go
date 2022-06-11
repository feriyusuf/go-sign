package main

import (
	"github.com/feriyusuf/go-sign/app/helpers"
	"time"
)

import (
	"github.com/feriyusuf/go-sign/app/controllers"
	"github.com/feriyusuf/go-sign/app/models_pg"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf(".env tidak ditemukan")
	}
}

func goDotEnvVariable(key string) string {
	return os.Getenv(key)
}

func main() {
	router := gin.Default()

	// Set local time zone from .env
	localTime, err := time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		panic("Can't set time zone")
	}
	time.Local = localTime

	// Connect to postgres
	models_pg.ConnectDatabase()

	// API Versioning
	withoutAuthorization := router.Group("/api/v1")
	withAuthorization := router.Group("/api/v1", helpers.TokenAuthentication(models_pg.PGDB))

	{
		auth := new(controllers.AuthController)
		withoutAuthorization.POST("/auth/register", auth.Register)
		withoutAuthorization.POST("/auth/login", auth.Login)

		// With middleware
		withAuthorization.DELETE("/auth/logout", auth.Logout)
		withAuthorization.GET("/auth/refresh", auth.Refresh)
	}

	// Default handle for unknown url address
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"Message": "Not found!"})
	})

	apiPort := ":" + goDotEnvVariable("API_PORT")

	router.Run(apiPort)
}
