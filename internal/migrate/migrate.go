package main

import (
	"fmt"

	"github.com/coderj001/smart_city_traffic_management_system/internal/models"
	"github.com/coderj001/smart_city_traffic_management_system/internal/repository"
)

func init() {
	repository.ConnectDB()
}

func main() {
	repository.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	repository.DB.AutoMigrate(&models.Sensor{}, &models.SensorData{})
	fmt.Println("👍 Migration complete")
}
