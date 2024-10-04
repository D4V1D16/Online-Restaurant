package controllers

import (
	"log"
	"net/http"
	"userAuth/Users/internal/Database"
	"userAuth/Users/internal/Utilities"

	"userAuth/Users/internal/Database/Models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context){

	var requestBody struct{
		Email string `json:"email"`
		Password string `json:"password"`
	}
	

	var user Models.User


	if err := c.ShouldBindJSON(&requestBody); err != nil {

        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	
	err := Database.DB.Where("email = ?", requestBody.Email).First(&user).Error
	
	if err != nil {
		log.Println("Error:", err)
	} else {

		log.Println("Password:", user.Password)
	}


	if !Utilities.CheckPassword(user.Password, requestBody.Password) {
		c.JSON(500, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	// Generate JWT
	jwt,err:=Utilities.GenerateToken(user)

	if err != nil {
		c.JSON(500, gin.H{
			"message":"Error generating the JWT",
			"error": err.Error(),
		})
		return
	}


	c.JSON(200, gin.H{
		"email": user.Email,
		"passwordFromDB": user.Password,
		"HashedPasswordFromRequest": requestBody.Password,
		"jwt": jwt,
	})

}

//This function will verify that the token is valid
func ProtectedRoute(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(401, gin.H{"error": "Missing or invalid token"})
		c.Abort()
		return
	}

	err := Utilities.VerifyToken(tokenString)

	if err != nil {
		c.JSON(401, gin.H{
			"error":   "Invalid JWT Token",
			"details": err.Error(),
		})
		c.Abort()
		return
	}

	// This part will extract the expire date of the token
	expClaim, err := Utilities.ExtractClaim(tokenString, "exp")

	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error extracting claim",
		})
		return
	}

	// Convert expClaim to float64 first, then to int64
	expFloat, ok := expClaim.(float64)
	if !ok {
		c.JSON(500, gin.H{
			"error": "Invalid claim type",
		})
		return
	}
	expTime := int64(expFloat)

	expired := Utilities.TokenExpired(expTime)

	if expired {
		c.JSON(401, gin.H{
			"error":   "Invalid JWT Token",
			"details": "Token expired",
		})
		return
	}

	// If everything is ok we will pass the route
	c.JSON(200, gin.H{
		"message": "This is a protected route",
		"details": "This token is valid",
	})
}
//This function will revoke the tokens
/*
func Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")



}
*/