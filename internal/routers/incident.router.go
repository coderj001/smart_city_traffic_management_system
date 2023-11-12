package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/coderj001/smart_city_traffic_management_system/internal/controllers"
)

type IncidentRouterController struct {
	incidentController controllers.IncidentController
}

func NewIncidentRouterController(incidentController controllers.IncidentController) *IncidentRouterController {
	return &IncidentRouterController{incidentController: incidentController}
}

func (i *IncidentRouterController) InitIncidentRouter(rg *gin.RouterGroup) {
	r1 := rg.Group("incidents")
	r1.GET("/", i.incidentController.GetIncidents)
	r1.POST("/", i.incidentController.CreateIncident)
	r2 := rg.Group("incident")
	r2.GET("/:id", i.incidentController.GetIncident)
	r2.PUT("/:id", i.incidentController.UpdateIncident)
}
