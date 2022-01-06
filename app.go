package main

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

	// Connect to postgres
	models_pg.ConnectDatabase(os.Getenv("DB_USER_PG"), os.Getenv("DB_PASSWORD_PG"), os.Getenv("DB_PORT_PG"), os.Getenv("DB_HOST_PG"), os.Getenv("DB_NAME_PG"))

	// API Versioning
	v1 := router.Group("/api/v1")
	{
		auth := new(controllers.AuthController)
		v1.POST("/auth/register", auth.Register)
		v1.POST("/auth/login", auth.Login)
	}

	// Default handle for unknown url address
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"Message": "Not found!"})
	})

	apiPort := ":" + goDotEnvVariable("API_PORT")

	router.Run(apiPort)
}
