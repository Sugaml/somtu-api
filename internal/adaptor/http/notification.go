package http

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// AddNotification	godoc
// @Summary			Add a new Notifications
// @Description		Add a new Notifications
// @Tags			Notifications
// @Accept			json
// @Produce			json
// @Security 		ApiKeyAuth
// @Param			NotificationRequest			body		domain.NotificationRequest		true		"Add Notification Request"
// @Success			200					{object}			domain.NotificationResponse					"Notification created"
// @Router			/notifications				[post]
func (h *Handler) CreateNotification(ctx *gin.Context) {
	var req *domain.NotificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	result, err := h.svc.CreateNotification(req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// ListNotification godoc
// @Summary 		List Notification
// @Description 	List Notification
// @Tags 			Notifications
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			query 								query 		string 		false 	"query"
// @Success 		200 				{array} 		domain.NotificationResponse
// @Router 			/notifications	 	[get]
func (h *Handler) ListNotification(ctx *gin.Context) {
	var req domain.ListNotificationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	req.Prepare()
	result, count, err := h.svc.ListNotification(&req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result, WithPagination(count, req.Page, req.Size))
}

// GetNotification 	godoc
// @Summary 		Get Notification
// @Description 	Get Notifications from Id
// @Tags 			Notifications
// @Accept  		json
// @Produce  		json
// @Security 		ApiKeyAuth
// @Param 			id path string true "Notifications id"
// @Success 		200 {object} domain.NotificationResponse
// @Router 			/notifications/{id} [get]
func (h *Handler) GetNotification(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := h.svc.GetNotification(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}

// UpdateNotification	godoc
// @Summary 			Update Notification
// @Description 		Update Notification from Id
// @Tags 				Notifications
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 							path 		string 								true 	"Notifications id"
// @Param 				NotificationsUpdateRequest	 body 		domain.UpdateNotificationRequest 	true 	"Update Notifications Response request"
// @Success 			200 						{object} 	domain.NotificationResponse
// @Router 				/Notificationss/{id} 		[put]
func (h *Handler) UpdateNotification(ctx *gin.Context) {
	id := ctx.Param("id")
	var req *domain.UpdateNotificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	data, err := h.svc.UpdateNotification(id, req)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, data)
}

// UpdateNotification	godoc
// @Summary 			Update Notification
// @Description 		Update Notification from Id
// @Tags 				Notifications
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Param 				id 							path 		string 								true 	"Notifications id"
// @Param 				NotificationsUpdateRequest	 body 		domain.UpdateNotificationRequest 	true 	"Update Notifications Response request"
// @Success 			200 						{object} 	domain.NotificationResponse
// @Router 				/Notificationss/read-all 		[put]
func (h *Handler) ReadAllNotification(ctx *gin.Context) {
	data, err := h.svc.ReadAllNotification()
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, data)
}

// DeleteNotification 	godoc
// @Summary 			Delete Notification
// @Description 		Delete Notification from Id
// @Tags 				Notifications
// @Accept  			json
// @Produce  			json
// @Security 			ApiKeyAuth
// @Security 			NotificationsAuth
// @Param 				id 						path 		string 						true 	"Notifications id"
// @Success 			200 					{object} 	domain.NotificationResponse
// @Router 				/notifications/{id} 	[delete]
func (ch *Handler) DeleteNotification(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ErrorResponse(ctx, http.StatusBadRequest, errors.New("required notification id"))
		return
	}
	result, err := ch.svc.DeleteNotification(id)
	if err != nil {
		ErrorResponse(ctx, http.StatusBadRequest, err)
		return
	}
	SuccessResponse(ctx, result)
}
