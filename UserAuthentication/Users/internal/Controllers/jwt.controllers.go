package controllers

import(
	"github.com/gin-gonic/gin"
	"userAuth/Users/internal/Database"
	"userAuth/Users/internal/Database/Models"
)

func Login(c *gin.Context){

	

	email := c.Param("email")
	password := c.Param("password")

	err := Database.DB.Select("password").Where("email = ?", email).First(&Models.User{}).Error


	if err != nil {
		c.JSON(500, gin.H{
			"error": "User not found",
		})
		return
	}



	c.JSON(200, gin.H{
		"email": email,
		"password": password,
	})

}