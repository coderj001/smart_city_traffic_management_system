package services

import "github.com/gin-gonic/gin"

// Handles gin http errors
func HandleError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(
		statusCode,
		gin.H{"error": err.Error()},
	)
	return
}
