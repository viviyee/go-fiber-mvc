package appsetting

import (
	"log"
	"os"

	"github.com/viviyee/go-mvc/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("MYSQL_DSN")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Post{})
}
