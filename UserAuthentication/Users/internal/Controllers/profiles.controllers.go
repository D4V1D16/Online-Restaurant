package controllers

import (
	"net/http"
	"userAuth/Users/internal/Database"
	"userAuth/Users/internal/Database/Models"

	"github.com/gin-gonic/gin"
)

func GetAllProfiles(c *gin.Context) {
	var profiles []Models.Profile
	Database.DB.Find(&profiles)

	c.JSON(200, gin.H{
		"Profiles": profiles,
	})

}


func CreateProfile(c *gin.Context) {
	var Profile Models.Profile


	if err := c.ShouldBindJSON(&Profile); err != nil {

        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


	result := Database.DB.Create(&Profile)


	if result.Error != nil {
		c.JSON(500, gin.H{
			"description":"Error creating in the database",
			"error": result.Error.Error(),
		})
		return
	}
	

	c.JSON(200, gin.H{
		"Status": "OK",
		"user": Profile.IdProfile,
		"Created": Profile.CreatedAt,
	})

}



func GetSingleProfile(c *gin.Context) {
	id := c.Param("id")

	var profile Models.Profile
	Database.DB.Where("id = ?", id).First(&profile)

    if profile.IdProfile == 0 {
        c.JSON(500, gin.H{
            "error": "Profile not found",
        })
        return
    }

    c.JSON(200, gin.H{
        "user": profile,
    })	
}

func UpdateProfile(c *gin.Context) {

	id := c.Param("id")

	exists := Database.DB.Where("id = ?", id).First(&Models.Profile{})

	if exists.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"description":"Profile does not exist",
			"error": "Profile not found",
		})
		return
	}


	var profile Models.Profile

	if err:= c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := Database.DB.Model(&Models.User{}).Where("id = ?", id).Omit("IdUser","ProfileID").Updates(&profile)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"description":"Error updating in the database",
			"error": result.Error.Error(),
		})
		return
	}


	c.JSON(200, gin.H{
		"Status": "OK",
		"Updated": profile.UpdatedAt,
	})
}

func DeleteProfile(c *gin.Context) {

	id := c.Param("id")

	exists := Database.DB.Where("id = ?", id).First(&Models.Profile{})

	if exists.RowsAffected == 0 {
		c.JSON(500, gin.H{
			"description":"User does not exist",
			"error": "User not found",
		})
		return
	}

	result := Database.DB.Delete(&Models.Profile{}, id)



	if result.Error != nil {
		c.JSON(500, gin.H{
			"description":"Error deleting in the database",
			"error": result.Error.Error(),
		})
		return
	}


	c.JSON(200, gin.H{
		"Status": "OK",
		"Deleted": id,
	})
}
