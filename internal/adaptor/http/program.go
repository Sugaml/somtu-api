package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// Program godoc
// @Summary			Create a new Program
// @Description		create a new Program
// @Tags			Program
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			createProgramRequest	body		domain.ProgramRequest	true	"Create Program Response request"
// @Success			200						{object}	domain.ProgramResponse			"Program Response"
// @Router			/programs [post]
func (ch *Handler) CreateProgram(ctx *gin.Context) {
	var req domain.ProgramRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	Program, err := ch.svc.CreateProgram(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, Program)
}

// ListProgram godoc
// @Summary List Program
// @Description List Categories
// @Tags Program
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param size query string false "Size"
// @Param page query string false "Page"
// @Param search query string false "Search"
// @Param sort-column query string false "Sort-Column"
// @Param sort-Direction query string false "Sort-Direction"
// @Success 200 {array} domain.ProgramResponse
// @Router /programs [get]
func (ch *Handler) ListProgram(ctx *gin.Context) {
	var req domain.ListProgramRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	result, count, err := ch.svc.LisProgram(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// GetProgram godoc
// @Summary Get Program
// @Description Get Program from Id
// @Tags Program
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "Program id"
// @Success 200 {object} domain.ProgramResponse
// @Router /programs/{id} [get]
func (ch *Handler) GetProgram(ctx *gin.Context) {
	id := ctx.Param("id")
	Program, err := ch.svc.GetProgram(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, Program)
}

// UpdateProgram 		godoc
// @Summary 			Update Program
// @Description 		Update Program from Id
// @Tags 				Program
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 								path 		string 							true 	"Program id"
// @Param 				ProgramUpdateRequest 			body 		domain.ProgramUpdateRequest 	true 	"Update Program Response request"
// @Success 			200 							{object} 	domain.ProgramResponse
// @Router 				/programs/{id} 				[put]
func (ch *Handler) UpdateProgram(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.ProgramUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	_, err := ch.svc.GetProgram(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	Program, err := ch.svc.UpdateProgram(ctx, id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, Program)
}

// DeleteProgram 		godoc
// @Summary 			Delete Program
// @Description 		Delete Program from Id
// @Tags 				Program
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 					path 		string 			true 	"Program id"
// @Success 			200 				{object} 	domain.ProgramResponse
// @Router 				/programs/{id} 	[delete]
func (ch *Handler) DeleteProgram(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required program id"))
		return
	}
	Program, err := ch.svc.Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = ch.svc.DeleteProgram(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, Program)
}
