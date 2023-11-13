package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/coderj001/smart_city_traffic_management_system/internal/config"
	"github.com/coderj001/smart_city_traffic_management_system/internal/controllers"
	"github.com/coderj001/smart_city_traffic_management_system/internal/repository"
	"github.com/coderj001/smart_city_traffic_management_system/internal/routers"
)

var (
	SensorController           controllers.SensorController
	SensorRouterController     routers.SensorRouterController
	SensorDataController       controllers.SensorDataController
	SensorDataRouterController routers.SensorDataRouterController
	IncidentController         controllers.IncidentController
	IncidentRouterController   routers.IncidentRouterController
	server                     *gin.Engine
)

func init() {
	repository.ConnectDB()
  repository.ConnectRedis()

	// routes
	// sensors
	SensorController = *controllers.NewSensorController(repository.DB)
	SensorRouterController = *routers.NewSensorRouterController(SensorController)

	// sensor data - traffic data
	SensorDataController = *controllers.NewSensorDataController(repository.DB)
	SensorDataRouterController = *routers.NewSensorDataRouterController(SensorDataController)

	// incidents
	IncidentController = *controllers.NewIncidentController(repository.DB)
	IncidentRouterController = *routers.NewIncidentRouterController(IncidentController)

	server = gin.Default()
}

func main() {
	r := server.Group("/")
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "🚦Smart Traffic Management System is running.")
	})

	router := server.Group("/api/v1")
	SensorRouterController.InitSensorRouter(router)
	SensorDataRouterController.InitSensorDataRouter(router)
	IncidentRouterController.InitIncidentRouter(router)

	log.Fatal(server.Run(":" + config.GetConfig().Port))
}
