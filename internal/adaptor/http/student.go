package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// AddStudent		godoc
// @Summary			Add a new Student
// @Description		Add a new Student
// @Tags			Student
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			StudentRequest			body		domain.UserRequest		true		"Add Student Request"
// @Success			200						{object}	domain.StudentResponse				"Student created"
// @Router			/students 				[post]
func (h *Handler) CreateStudent(ctx *gin.Context) {
	var req *domain.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateUser(req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// AddStudent		godoc
// @Summary			Add a new Student
// @Description		Add a new Student
// @Tags			Student
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			StudentRequest			body		domain.UserRequest		true		"Add Student Request"
// @Success			200						{object}	domain.StudentResponse				"Student created"
// @Router			/students/bulk 			[post]
func (h *Handler) CreateBulkStudent(ctx *gin.Context) {
	var req *[]domain.UserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateBulkUser(req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
