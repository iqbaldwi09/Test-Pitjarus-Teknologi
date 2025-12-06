package handler

import (
	"net/http"
	"time"

	"go-backend/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	uc usecase.ReportUsecase
}

func NewReportHandler(uc usecase.ReportUsecase) *ReportHandler {
	return &ReportHandler{uc}
}

func (h *ReportHandler) GetReports(c *gin.Context) {

	dateFromStr := c.Query("dateFrom")
	dateToStr := c.Query("dateTo")
	areaID := c.Query("area_id")

	var dateFrom, dateTo time.Time
	var err error

	if dateFromStr != "" && dateToStr != "" {
		dateFrom, err = time.Parse("2006-01-02", dateFromStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dateFrom format, use YYYY-MM-DD"})
			return
		}

		dateTo, err = time.Parse("2006-01-02", dateToStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dateTo format, use YYYY-MM-DD"})
			return
		}
	}

	chart, errChart := h.uc.GetAreaChart(dateFrom, dateTo, areaID)
	table, errTable := h.uc.GetBrandTable(dateFrom, dateTo, areaID)

	if errChart != nil || errTable != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch report data",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chart": chart,
		"table": table,
	})
}

func (h *ReportHandler) GetAreas(c *gin.Context) {
	areas, err := h.uc.GetAllAreas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch areas",
		})
		return
	}

	c.JSON(http.StatusOK, areas)
}
