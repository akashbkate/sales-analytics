package handler

import (
	service "sales-analytics/internals/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/api/revenue", service.GetTotalRevenue)
	r.POST("/api/refresh", service.RefreshData)
}
