package router

import (
	"go-backend/internal/handler"
	"go-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(reportHandler *handler.ReportHandler) *gin.Engine {
	r := gin.Default()

	// middleware CORS dan API key
	r.Use(middleware.CORS())
	r.Use(middleware.APIKeyAuth())

	api := r.Group("/api")
	{
		api.GET("/reports", reportHandler.GetReports)
		api.GET("/areas", reportHandler.GetAreas)
	}

	return r
}
