package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// AddFine			godoc
// @Summary			Add a new Fine
// @Description		Add a new Fine
// @Tags			Fine
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			FineRequest			body		domain.FineRequest		true		"Add Fine Request"
// @Success			200					{object}	domain.FineResponse					"Fine created"
// @Router			/fines 				[post]
func (h *Handler) CreateFine(ctx *gin.Context) {
	var req *domain.FineRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateFine(req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListFine 		godoc
// @Summary 		List Fine
// @Description 	List Fine
// @Tags 			Fine
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			query 						query 		string 		false 	"query"
// @Success 		200 		{array} 		domain.FineResponse
// @Router 			/fines	 	[get]
func (h *Handler) ListFine(ctx *gin.Context) {
	var req domain.ListFineRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	result, count, err := h.svc.ListFine(&req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// GetFine 			godoc
// @Summary 		Get Fine
// @Description 	Get Fine from Id
// @Tags 			Fine
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "Fine id"
// @Success 		200 {object} domain.FineResponse
// @Router 			/fines/{id} [get]
func (h *Handler) GetFine(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.GetFine(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateFine			godoc
// @Summary 			Update Fine
// @Description 		Update Fine from Id
// @Tags 				Fine
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 							path 		string 								true 	"Fine id"
// @Param 				FineUpdateRequest	 		body 		domain.UpdateFineRequest 	true 	"Update Fine Response request"
// @Success 			200 						{object} 	domain.FineResponse
// @Router 				/fines/{id} 				[put]
func (h *Handler) UpdateFine(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.UpdateFineRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := h.svc.UpdateFine(id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, data)
}

// DeleteFine 			godoc
// @Summary 			Delete Fine
// @Description 		Delete Fine from Id
// @Tags 				Fine
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Security 			FineAuth
// @Param 				id 						path 		string 						true 	"Fine id"
// @Success 			200 					{object} 	domain.FineResponse
// @Router 				/Fines/{id} 	[delete]
func (ch *Handler) DeleteFine(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required Fine id"))
		return
	}
	result, err := ch.svc.DeleteFine(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
