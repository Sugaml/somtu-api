package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	storage "github.com/sugaml/lms-api/internal/adaptor/storage/uploader"
)

func (h *Handler) UploadFile(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("file required"))
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("file open failed"))
		return
	}
	defer file.Close()

	// Pass to your uploader
	fileDetails := &storage.FileDetails{
		FileType: storage.FileTypeStudentPhoto,
		// EntityID:   r.FormValue("student_id"),
		File:       file,
		FileHeader: fileHeader,
		Metadata: storage.FileMetadata{
			FileName:    fileHeader.Filename,
			Size:        fileHeader.Size,
			ContentType: fileHeader.Header.Get("Content-Type"),
			UploadedAt:  time.Now(),
		},
	}
	result, err := h.uploader.UploadFile(fileDetails)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, map[string]string{
		"url": result,
	})
}
