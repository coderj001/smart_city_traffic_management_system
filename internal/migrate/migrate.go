package main

import (
	"fmt"

	"github.com/coderj001/smart_city_traffic_management_system/internal/models"
	"github.com/coderj001/smart_city_traffic_management_system/internal/repository"
	"github.com/coderj001/smart_city_traffic_management_system/internal/services"
)

func init() {
	repository.ConnectDB()
	repository.ConnectRedis()
}

func main() {
	repository.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	repository.DB.AutoMigrate(
		&models.Sensor{},
		&models.SensorData{},
		&models.Incident{},
	)
	fmt.Println("👍 Migration complete")
	data := map[string]interface{}{
		"field1": "value1",
		"field2": 123,
	}
	services.Set("data", data, 3600)
  services.Exists("data")
}
