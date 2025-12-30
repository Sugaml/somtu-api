package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// AddBook			godoc
// @Summary			Add a new Book
// @Description		Add a new Book
// @Tags			Book
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			BookRequest			body		domain.BookRequest		true		"Add Book Request"
// @Success			200					{object}	domain.BookResponse					"Book created"
// @Router			/books 				[post]
func (h *Handler) CreateBook(ctx *gin.Context) {
	user_id, exists := ctx.Get(authorizationUserrIDKey)
	if !exists {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("authorization user id not found"))
		return
	}
	logrus.Info("Authorization user id: ", user_id)
	var req *domain.BookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateBook(ctx, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListBook 		godoc
// @Summary 		List Book
// @Description 	List Book
// @Tags 			Book
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			query 						query 		string 		false 	"query"
// @Success 		200 		{array} 		domain.BookResponse
// @Router 			/books	 	[get]
func (h *Handler) ListBook(ctx *gin.Context) {
	var req domain.BookListRequest
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
	result, count, err := h.svc.ListBook(ctx, &req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// GetBook 			godoc
// @Summary 		Get Book
// @Description 	Get Book from Id
// @Tags 			Book
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "Book id"
// @Success 		200 {object} domain.BookResponse
// @Router 			/books/{id} [get]
func (h *Handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.GetBook(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateBook		godoc
// @Summary 			Update Book
// @Description 		Update Book from Id
// @Tags 				Book
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 							path 		string 								true 	"Book id"
// @Param 				BookUpdateRequest	 		body 		domain.BookUpdateRequest 	true 	"Update Book Response request"
// @Success 			200 						{object} 	domain.BookResponse
// @Router 				/books/{id} 				[put]
func (h *Handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.BookUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := h.svc.UpdateBook(ctx, id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, data)
}

// DeleteBook 			godoc
// @Summary 			Delete Book
// @Description 		Delete Book from Id
// @Tags 				Book
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Security 			BookAuth
// @Param 				id 						path 		string 						true 	"Book id"
// @Success 			200 					{object} 	domain.BookResponse
// @Router 				/books/{id} 	[delete]
func (ch *Handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required Book id"))
		return
	}
	result, err := ch.svc.DeleteBook(ctx, id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
