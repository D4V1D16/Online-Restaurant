package Database

import (
	"os"
	"log"


	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading de enviroment files in databas connection")
	}
}


var DB *gorm.DB

func DatabaseConn() {
    dsn := os.Getenv("DB_URL")
    var err error

    
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }

    log.Println("Database connected")
}