package main

import (
	"net/http"

	"github.com/coderj001/smart_city_traffic_management_system/internal/repository"
	"github.com/gin-gonic/gin"
)

func init() {
	repository.ConnectDB()
}

func main() {
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "🚦Smart Traffic Management System is running.")
	})
	r.Run(":8000")
}
