package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// response represents a response body format
type responseData struct {
	Error   int    `json:"error" example:"0"`
	Message string `json:"message" example:"Message"`
	Data    any    `json:"data,omitempty"`
}

type metaData struct {
	Count int64 `json:"count,omitempty"`
	Page  int   `json:"page,omitempty"`
	Size  int   `json:"size,omitempty"`
}

type response struct {
	responseData
	metaData
}

type metaOptions struct {
	Error   int    `json:"error"`
	Message string `json:"message"`
	Count   int64  `json:"count"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
}

type SuccessOption func(req *metaOptions)

func WithPagination(count int64, page, size int) SuccessOption {
	return func(req *metaOptions) {
		req.Count = count
		req.Page = page
		req.Size = size
	}
}

func WithMessage(message string) SuccessOption {
	return func(req *metaOptions) {
		req.Message = message
	}
}
func WithError(error int) SuccessOption {
	return func(req *metaOptions) {
		req.Error = error
	}
}

func SuccessResponse(ctx *gin.Context, data any, opts ...SuccessOption) {
	res := &metaOptions{}
	for _, opt := range opts {
		opt(res)
	}
	ctx.JSON(http.StatusOK, response{
		responseData: responseData{
			Error:   res.Error,
			Message: res.Message,
			Data:    data,
		},
		metaData: metaData{
			Count: res.Count,
			Page:  res.Page,
			Size:  res.Size,
		},
	})
}

func ErrorResponse(ctx *gin.Context, code int, err error) {
	logrus.Errorf("Error response logged : %v", err)
	ctx.JSON(code, responseData{
		Error:   code,
		Message: err.Error(),
	})
}
