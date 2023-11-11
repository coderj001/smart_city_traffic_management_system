package routers

import (
	"github.com/coderj001/smart_city_traffic_management_system/internal/controllers"
	"github.com/gin-gonic/gin"
)

type SensorDataRouterController struct {
	sensorDataController controllers.SensorDataController
}

func NewSensorDataRouterController(sensorDataController controllers.SensorDataController) *SensorDataRouterController {
	return &SensorDataRouterController{sensorDataController: sensorDataController}
}

func (s *SensorDataRouterController) InitSensorDataRouter(rg *gin.RouterGroup) {
	router := rg.Group("traffic")
  router.POST("/", s.sensorDataController.PostSensorData)
	router.GET("/:id", s.sensorDataController.GetSensorData)
}
