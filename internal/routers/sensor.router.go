package routers

import (
	"github.com/coderj001/smart_city_traffic_management_system/internal/controllers"
	"github.com/gin-gonic/gin"
)

type SensorRouterController struct {
	sensorController controllers.SensorController
}

func NewSensorRouterController(sensorController controllers.SensorController) *SensorRouterController {
	return &SensorRouterController{sensorController: sensorController}
}

func (s *SensorRouterController) InitSensorRouter(rg *gin.RouterGroup) {
	router := rg.Group("sensor")
	router.GET("/", s.sensorController.GetAll)
	router.POST("/", s.sensorController.Create)
	router.GET("/:id", s.sensorController.GetByID)
	router.PUT("/:id", s.sensorController.UpdateByID)
}
