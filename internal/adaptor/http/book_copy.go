package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// AddBookCopy		godoc
// @Summary			Add a new BookCopy
// @Description		Add a new BookCopy
// @Tags			BookCopy
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			BookCopyRequest			body		domain.BookCopyRequest		true		"Add BookCopy Request"
// @Success			200					{object}	domain.BookCopyResponse					"BookCopy created"
// @Router			/book-copies			[post]
func (h *Handler) CreateBookCopy(ctx *gin.Context) {
	user_id, exists := ctx.Get(authorizationUserrIDKey)
	if !exists {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("authorization user id not found"))
		return
	}
	logrus.Info("Authorization user id: ", user_id)
	var req *domain.AddBookCopiesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateBookCopy(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListBookCopy 		godoc
// @Summary 		List BookCopy
// @Description 	List BookCopy
// @Tags 			BookCopy
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			query 						query 		string 		false 	"query"
// @Success 		200 		{array} 		domain.BookCopyResponse
// @Router 			/book-copies	 	[get]
func (h *Handler) ListBookCopy(ctx *gin.Context) {
	var req domain.BookCopyListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	user_id, exists := ctx.Get(authorizationUserrIDKey)
	if !exists {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("authorization user id not found"))
		return
	}
	logrus.Info("Authorization user id: ", user_id)
	if req.Status == "" {
		req.Status = "available"
	}
	result, count, err := h.svc.ListBookCopies(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// ListBookCopy 		godoc
// @Summary 		List BookCopy
// @Description 	List BookCopy
// @Tags 			BookCopy
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			query 						query 		string 		false 	"query"
// @Success 		200 		{array} 		domain.BookCopyResponse
// @Router 			/books/:id/book-copies		[get]
func (h *Handler) ListBookCopyByBookId(ctx *gin.Context) {
	id := ctx.Param("id")
	var req domain.BookCopyListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	user_id, exists := ctx.Get(authorizationUserrIDKey)
	if !exists {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("authorization user id not found"))
		return
	}
	logrus.Info("Authorization user id: ", user_id)
	if req.Status == "" {
		req.Status = "available"
	}
	result, count, err := h.svc.ListBookCopiesByBookId(ctx, id, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// GetBookCopy 			godoc
// @Summary 		Get BookCopy
// @Description 	Get BookCopy from Id
// @Tags 			BookCopy
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "BookCopy id"
// @Success 		200 {object} domain.BookCopyResponse
// @Router 			/book-copies/{id} [get]
func (h *Handler) GetBookCopy(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.GetBookCopy(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateBookCopy		godoc
// @Summary 			Update BookCopy
// @Description 		Update BookCopy from Id
// @Tags 				BookCopy
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 							path 		string 								true 	"BookCopy id"
// @Param 				BookCopyUpdateRequest	 		body 		domain.BookCopyUpdateRequest 	true 	"Update BookCopy Response request"
// @Success 			200 						{object} 	domain.BookCopyResponse
// @Router 				/book-copies/{id} 				[put]
func (h *Handler) UpdateBookCopy(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.BookCopyUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := h.svc.UpdateBookCopy(ctx, id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, data)
}

// DeleteBookCopy 		godoc
// @Summary 			Delete BookCopy
// @Description 		Delete BookCopy from Id
// @Tags 				BookCopy
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Security 			BookCopyAuth
// @Param 				id 						path 		string 						true 	"BookCopy id"
// @Success 			200 					{object} 	domain.BookCopyResponse
// @Router 				/book-copies/{id} 	[delete]
func (ch *Handler) DeleteBookCopy(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required BookCopy id"))
		return
	}
	result, err := ch.svc.DeleteBookCopy(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
