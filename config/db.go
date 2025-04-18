package config

import (
	"log"
	"os"

	"github.com/Wasay1567/edutrack/models"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	godotenv.Load()
	dsn := os.Getenv("DB_CONNECTION_STRING")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed")
	}
	err = DB.AutoMigrate(&models.Student{}, &models.Course{}, &models.Enrollment{})
	if err != nil {
		log.Fatal("Migration failed")
	}
	log.Println("Database connected successfully")
}
