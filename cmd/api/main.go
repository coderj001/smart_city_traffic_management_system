package main

import (
	"log"
	"net/http"

	"github.com/coderj001/smart_city_traffic_management_system/internal/initilizers"
	"github.com/gin-gonic/gin"
)

func init() {
	config, err := initilizers.LoadConfig()
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}
	initilizers.ConnectDB(&config)
}

func main() {
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Smart Traffic Management System is running.")
	})
	r.Run(":8000")
}
