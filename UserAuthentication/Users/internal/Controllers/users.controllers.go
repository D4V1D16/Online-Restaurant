package controllers

import (
	"net/http"
	"userAuth/Users/internal/Database"
	"userAuth/Users/internal/Database/Models"
	"userAuth/Users/internal/Utilities"

	"github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
	var users []Models.User
	Database.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})

}


func CreateUser(c *gin.Context) {
	var User Models.User


	if err := c.ShouldBindJSON(&User); err != nil {

        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


	emailValidation := Database.DB.Where("email = ?", User.Email).First(&User)

	if emailValidation.RowsAffected > 0 {
		c.JSON(500, gin.H{
			"error": "Email already exists",
		})
		return
	}

	idValidation := Database.DB.Where("id_user = ?", User.IdUser).First(&User)

	if idValidation.RowsAffected > 0 {
		c.JSON(500, gin.H{
			"error": "ID already exists",
		})
		return
	}

	profileValidation := Database.DB.Where("id_profile = ?", User.ProfileID).First(&User)

	if profileValidation.RowsAffected > 0 {
		c.JSON(500, gin.H{
			"error": "Profile already exists",
		})
		return
	}


	User.Password,_ = Utilities.HashPassword(User.Password)










	result := Database.DB.Create(&User)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"description":"Error creating in the database",
			"error": result.Error.Error(),
		})
		return
	}
	

	c.JSON(200, gin.H{
		"Status": "OK",
		"user": User.IdUser,
		"Created": User.CreatedAt,
		"HashedPassword": User.Password,
	})



}



func GetSingleUser(c *gin.Context) {
	id := c.Param("id")

	var user Models.User
	Database.DB.Where("id = ?", id).First(&user)

    if user.IdUser == 0 {
        c.JSON(500, gin.H{
            "error": "User not found",
        })
        return
    }

    c.JSON(200, gin.H{
        "user": user,
    })	
}

func UpdateUser() {

}

func DeleteUser(c *gin.Context) {

	id := c.Param("id")

	exists := Database.DB.Where("id = ?", id).First(&Models.User{})

	if exists.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"error": "User not found",
		})
		return
	}

	result := Database.DB.Delete(&Models.User{}, id)



	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}


	c.JSON(200, gin.H{
		"Status": "OK",
		"Deleted": id,
	})
}
