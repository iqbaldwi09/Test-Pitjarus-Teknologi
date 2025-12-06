package repository

import (
	"time"

	"gorm.io/gorm"
)

type ReportRepository interface {
	GetAreaChart(dateFrom, dateTo time.Time, areaID string) ([]AreaChartResult, error)
	GetBrandTable(dateFrom, dateTo time.Time, areaID string) ([]BrandTableResult, error)
	GetAllAreas() ([]Area, error)
}

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{db: db}
}

type AreaChartResult struct {
	AreaName string  `json:"area"`
	Value    float64 `json:"value"`
}

func (r *reportRepository) GetAreaChart(dateFrom, dateTo time.Time, areaID string) ([]AreaChartResult, error) {
	var result []AreaChartResult

	query := r.db.
		Table("report_product AS rp").
		Select(`
			sa.area_name AS area_name,
			(SUM(rp.compliance) / COUNT(*) * 100) AS value
		`).
		Joins("JOIN store s ON s.store_id = rp.store_id").
		Joins("JOIN store_area sa ON sa.area_id = s.area_id")

	if !dateFrom.IsZero() && !dateTo.IsZero() {
		query = query.Where("rp.tanggal BETWEEN ? AND ?", dateFrom, dateTo)
	}

	if areaID != "" {
		query = query.Where("sa.area_id = ?", areaID)
	}

	err := query.
		Group("sa.area_id, sa.area_name").
		Scan(&result).Error

	return result, err
}

type BrandTableResult struct {
	BrandName string  `json:"brand"`
	AreaName  string  `json:"area"`
	Value     float64 `json:"value"`
}

func (r *reportRepository) GetBrandTable(dateFrom, dateTo time.Time, areaID string) ([]BrandTableResult, error) {
	var result []BrandTableResult

	query := r.db.
		Table("report_product AS rp").
		Select(`
			b.brand_name AS brand_name,
			sa.area_name AS area_name,
			(SUM(rp.compliance) / COUNT(*) * 100) AS value
		`).
		Joins("JOIN product p ON p.product_id = rp.product_id").
		Joins("JOIN product_brand b ON b.brand_id = p.brand_id").
		Joins("JOIN store s ON s.store_id = rp.store_id").
		Joins("JOIN store_area sa ON sa.area_id = s.area_id")

	if !dateFrom.IsZero() && !dateTo.IsZero() {
		query = query.Where("rp.tanggal BETWEEN ? AND ?", dateFrom, dateTo)
	}

	if areaID != "" {
		query = query.Where("sa.area_id = ?", areaID)
	}

	err := query.
		Group("b.brand_name, sa.area_name").
		Scan(&result).Error

	return result, err
}

type Area struct {
	ID   string `json:"id" gorm:"column:area_id"`
	Name string `json:"name" gorm:"column:area_name"`
}

func (r *reportRepository) GetAllAreas() ([]Area, error) {
	var areas []Area

	err := r.db.
		Table("store_area").
		Select("area_id, area_name").
		Order("area_name ASC").
		Scan(&areas).Error

	return areas, err
}
