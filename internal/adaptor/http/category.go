package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// Category godoc
// @Summary			Create a new Category
// @Description		create a new Category
// @Tags			Category
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			createCategoryRequest	body		domain.CategoryRequest	true	"Create Category Response request"
// @Success			200						{object}	domain.CategoryResponse			"Category Response"
// @Router			/categories [post]
func (ch *Handler) CreateCategory(ctx *gin.Context) {
	var req domain.CategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	category, err := ch.svc.Create(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, category)
}

// ListCategory godoc
// @Summary List Category
// @Description List Categories
// @Tags Category
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param size query string false "Size"
// @Param page query string false "Page"
// @Param search query string false "Search"
// @Param sort-column query string false "Sort-Column"
// @Param sort-Direction query string false "Sort-Direction"
// @Success 200 {array} domain.CategoryResponse
// @Router /categories [get]
func (ch *Handler) ListCategory(ctx *gin.Context) {
	var req domain.ListCategoryRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	result, count, err := ch.svc.List(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// GetCategory godoc
// @Summary Get Category
// @Description Get Category from Id
// @Tags Category
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path string true "category id"
// @Success 200 {object} domain.CategoryResponse
// @Router /categories/{id} [get]
func (ch *Handler) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	category, err := ch.svc.Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, category)
}

// UpdateCategory 		godoc
// @Summary 			Update Category
// @Description 		Update Category from Id
// @Tags 				Category
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 								path 		string 							true 	"category id"
// @Param 				CategoryUpdateRequest 			body 		domain.CategoryUpdateRequest 	true 	"Update Category Response request"
// @Success 			200 							{object} 	domain.CategoryResponse
// @Router 				/categories/{id} 				[put]
func (ch *Handler) UpdateCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.CategoryUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	_, err := ch.svc.Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	category, err := ch.svc.Update(ctx, id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, category)
}

// DeleteCategory 		godoc
// @Summary 			Delete Category
// @Description 		Delete Category from Id
// @Tags 				Category
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 					path 		string 			true 	"category id"
// @Success 			200 				{object} 	domain.CategoryResponse
// @Router 				/categories/{id} 	[delete]
func (ch *Handler) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required category id"))
		return
	}
	category, err := ch.svc.Get(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	err = ch.svc.Delete(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	SuccessResponse(ctx, category)
}
