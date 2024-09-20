package main

import (
	"log"
	"net/http"
	"userAuth/Users/internal/Database"
	"userAuth/Users/internal/API"

	"gorm.io/gorm"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading de enviroment files in the main file" + err.Error())
	} else {
		Database.DatabaseConn()
		if err != nil {
			log.Fatal("Error connecting to the database" + err.Error())
		} else {
			log.Println("Database connected")
			Database.Automigrations()
		}

	}

}

func main() {

	router := gin.Default()

	API.UserRoutes(router)

	//Test Route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "I think the database is connected",
		})
	})

	router.Run()
}
