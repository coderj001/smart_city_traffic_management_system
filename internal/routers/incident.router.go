package routers

import (
	"github.com/coderj001/smart_city_traffic_management_system/internal/controllers"
	"github.com/gin-gonic/gin"
)

type IncidentRouterController struct {
	incidentController controllers.IncidentController
}

func NewIncidentRouterController(incidentController controllers.IncidentController) *IncidentRouterController {
	return &IncidentRouterController{incidentController: incidentController}
}

func (i *IncidentRouterController) InitIncidentRouter(rg *gin.RouterGroup) {
	router := rg.Group("incidents")
	router.GET("/", i.incidentController.GetIncidents)
  router.POST("/", i.incidentController.CreateIncident)
  router2 := rg.Group("incident")
  router2.GET("/:id", i.incidentController.GetIncident)
  router2.PUT("/:id", i.incidentController.UpdateIncident)
}
