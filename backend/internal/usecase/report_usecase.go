package usecase

import (
	"time"

	"go-backend/internal/repository"
)

type ReportUsecase interface {
	GetAreaChart(dateFrom, dateTo time.Time, areaID string) ([]repository.AreaChartResult, error)
	GetBrandTable(dateFrom, dateTo time.Time, areaID string) ([]repository.BrandTableResult, error)
	GetAllAreas() ([]repository.Area, error)
}

type reportUsecase struct {
	repo repository.ReportRepository
}

func NewReportUsecase(r repository.ReportRepository) ReportUsecase {
	return &reportUsecase{repo: r}
}

func (u *reportUsecase) GetAreaChart(dateFrom, dateTo time.Time, areaID string) ([]repository.AreaChartResult, error) {
	return u.repo.GetAreaChart(dateFrom, dateTo, areaID)
}

func (u *reportUsecase) GetBrandTable(dateFrom, dateTo time.Time, areaID string) ([]repository.BrandTableResult, error) {
	return u.repo.GetBrandTable(dateFrom, dateTo, areaID)
}

func (u *reportUsecase) GetAllAreas() ([]repository.Area, error) {
	return u.repo.GetAllAreas()
}
