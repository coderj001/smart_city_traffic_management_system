package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Smart Traffic Management System is running.")
	})
	r.Run(":8000")
}
