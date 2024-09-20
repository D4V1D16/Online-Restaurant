package Database

import (
	"log"
	"userAuth/Users/internal/Database/Models"

)

func Automigrations() {
	if DB == nil{
		log.Fatal("Database not connected")
	}

	DB.AutoMigrate(&Models.User{})
	DB.AutoMigrate(&Models.Profile{})
	DB.AutoMigrate(&Models.Role{})
	DB.AutoMigrate(&Models.Permission{})

	log.Println("Database migrated")

}
