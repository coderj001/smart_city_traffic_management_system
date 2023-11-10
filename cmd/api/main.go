package main

import (
	"log"
	"net/http"

	"github.com/coderj001/smart_city_traffic_management_system/internal/config"
	"github.com/coderj001/smart_city_traffic_management_system/internal/controllers"
	"github.com/coderj001/smart_city_traffic_management_system/internal/repository"
	"github.com/coderj001/smart_city_traffic_management_system/internal/routers"
	"github.com/gin-gonic/gin"
)

var (
	SensorController       controllers.SensorController
	SensorRouterController routers.SensorRouterController
	server                 *gin.Engine
)

func init() {
	repository.ConnectDB()

	// routes
	SensorController = *controllers.NewSensorController(repository.DB)
	SensorRouterController = *routers.NewSensorRouterController(SensorController)

	server = gin.Default()
}

func main() {
	r := server.Group("/")
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "🚦Smart Traffic Management System is running.")
	})

	router := server.Group("/api/v1")
	SensorRouterController.InitSensorRouter(router)

	log.Fatal(server.Run(":" + config.GetConfig().Port))
}
