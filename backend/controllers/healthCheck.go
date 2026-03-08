package controllers

import "github.com/gin-gonic/gin"

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(OK, gin.H{"message": "healthy"})
}
