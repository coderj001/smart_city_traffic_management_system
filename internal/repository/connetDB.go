package repository

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/coderj001/smart_city_traffic_management_system/internal/config"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dbConfig := config.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Kolkata",
		dbConfig.DB.Host,
		dbConfig.DB.Username,
		dbConfig.DB.Password,
		dbConfig.DB.Name,
		dbConfig.DB.Port,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
