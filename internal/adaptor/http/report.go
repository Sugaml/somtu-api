package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// ListLibraryDashboardStats	godoc
// @Summary 		List LibraryDashboard
// @Description 	List LibraryDashboard
// @Tags 			Report
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Success 		200 {object} domain.LibraryDashboardStats
// @Router 			/reports/dashboard-stats	[get]
func (h *Handler) GetLibraryDashboardStats(ctx *gin.Context) {
	result, err := h.svc.GetLibraryDashboardStats()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// GetMonthlyChartData	godoc
// @Summary 		Chart Data
// @Description 	Get chart data for various time ranges (daily, weekly, monthly, quarterly, yearly)
// @Tags 			Report
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			range 			query 		string 		false 	"Time range: daily, weekly, monthly, quarterly, yearly"
// @Param 			start_date 		query 		string 		false 	"Start date (YYYY-MM-DD)"
// @Param 			end_date 		query 		string 		false 	"End date (YYYY-MM-DD)"
// @Success 		200 {array} 			domain.ChartData
// @Router 			/reports/chart-stats	[get]
func (h *Handler) GetMonthlyChartData(ctx *gin.Context) {
	// today := time.Now()
	// sevenDaysAgo := today.AddDate(0, 0, -7) // inclusive
	// startDate := sevenDaysAgo.Format("2006-01-02")
	// endDate := today.Format("2006-01-02")
	var req domain.ChartRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.GetDailyChartData(&req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListBorrowStats	godoc
// @Summary 		List Borrow
// @Tags 			Report
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Success 		200 {object} domain.BorrowedBookStats
// @Router 			/reports/borrowedbookstats	[get]
func (h *Handler) GetBorrowedBookStats(ctx *gin.Context) {
	result, err := h.svc.GetBorrowedBookStats()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListBookProgramstats	godoc
// @Summary 			List Borrow
// @Tags 				Report
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Success 			200 {array} domain.BookProgramstats
// @Router 				/reports/program-stats	[get]
func (h *Handler) GetBookProgramstats(ctx *gin.Context) {
	result, err := h.svc.GetBookProgramstats()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// GetInventorystats	godoc
// @Summary 			List Borrow
// @Tags 				Report
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Success 			200 {array} domain.BookProgramstats
// @Router 				/reports/inventory-stats	[get]
func (h *Handler) GetInventorystats(ctx *gin.Context) {
	result, err := h.svc.GetInventorystats()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
