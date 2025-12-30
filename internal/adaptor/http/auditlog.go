package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// AddAuditLog	godoc
// @Summary			Add a new AuditLog
// @Description		Add a new AuditLog
// @Tags			AuditLog
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Security 		UserAuth
// @Param			AuditLogRequest			body			domain.AuditLogRequest		true		"Add Parking AuditLog Request"
// @Success			200						{object}		domain.AuditLogResponse					"Parking AuditLog created"
// @Router			/auditlog 		[post]
func (h *Handler) CreateAuditLog(ctx *gin.Context) {
	var req *domain.AuditLogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateAuditLog(req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListAuditLog 		godoc
// @Summary 		List AuditLog
// @Description 	List AuditLog from Id
// @Tags 			AuditLog
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Security 		UserAuth
// @Success 		200 {array} domain.AuditLogResponse
// @Router 			/auditlog [get]
func (h *Handler) ListAuditLog(ctx *gin.Context) {
	var req domain.ListAuditLogRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	result, count, err := h.svc.ListAuditLog(&req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, int(req.Page), int(req.Size)))
}

// GetAuditLog 		godoc
// @Summary 		Get AuditLog
// @Description 	Get AuditLog from Id
// @Tags 			AuditLog
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Security 		UserAuth
// @Param 			id path string true "AuditLog id"
// @Success 		200 {object} domain.AuditLogResponse
// @Router 			/auditlog/{id} [get]
func (h *Handler) GetAuditLog(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.GetAuditLog(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateAuditLog		godoc
// @Summary 			Update AuditLog
// @Description 		Update AuditLog from Id
// @Tags 				AuditLog
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Security 			UserAuth
// @Param 				id 								path 			string 								true 	"AuditLog id"
// @Param 				AuditLogUpdateRequest	 	body 			domain.AuditLogUpdateRequest 	true 	"Update AuditLog Response request"
// @Success 			200 							{object} 		domain.AuditLogResponse
// @Router 				/auditlog/{id} 				[put]
func (h *Handler) UpdateAuditLog(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.AuditLogUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := h.svc.UpdateAuditLog(id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, data)
}

// DeleteAuditLog 		godoc
// @Summary 			Delete AuditLog
// @Description 		Delete AuditLog from Id
// @Tags 				AuditLog
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Security 			UserAuth
// @Param 				id 						path 		string 						true 	"AuditLog id"
// @Success 			200 					{object} 	domain.AuditLogResponse
// @Router 				/auditlog/{id} 	[delete]
func (ch *Handler) DeleteAuditLog(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required audit log id"))
		return
	}
	err := ch.svc.DeleteAuditLog(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, nil)
}
