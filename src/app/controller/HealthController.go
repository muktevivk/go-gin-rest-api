package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	log "github.com/sirupsen/logrus"
)

//HealthCheck endpoint
func HealthCheck (context *gin.Context) {
	log.Debug("Health check")
	context.JSON(http.StatusOK, "Up!")
}