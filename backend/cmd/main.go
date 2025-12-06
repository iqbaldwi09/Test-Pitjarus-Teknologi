package main

import (
	"go-backend/config"
	"go-backend/internal/handler"
	"go-backend/internal/repository"
	"go-backend/internal/usecase"
	"go-backend/router"
)

func main() {
	db := config.ConnectDB()

	reportRepo := repository.NewReportRepository(db)
	reportUsecase := usecase.NewReportUsecase(reportRepo)
	reportHandler := handler.NewReportHandler(reportUsecase)

	r := router.SetupRouter(reportHandler)

	r.Run(":8080")
}
