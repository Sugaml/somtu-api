package port

import (
	"github.com/sugaml/lms-api/internal/core/domain"
)

// type NotificationRepository interface is an interface for interacting with type Announcement-related data
type NotificationRepository interface {
	CreateNotification(data *domain.Notification) (*domain.Notification, error)
	ListNotification(req *domain.ListNotificationRequest) ([]*domain.Notification, int64, error)
	GetNotification(id string) (*domain.Notification, error)
	UpdateNotification(id string, req domain.Map) (*domain.Notification, error)
	ReadAllNotification(req domain.Map) (*domain.Notification, error)
	DeleteNotification(id string) error
}

// type NotificationService interface is an interface for interacting with type Announcement-related data
type NotificationService interface {
	CreateNotification(data *domain.NotificationRequest) (*domain.NotificationResponse, error)
	ListNotification(req *domain.ListNotificationRequest) ([]*domain.NotificationResponse, int64, error)
	GetNotification(id string) (*domain.NotificationResponse, error)
	UpdateNotification(id string, req *domain.UpdateNotificationRequest) (*domain.NotificationResponse, error)
	ReadAllNotification() (*domain.NotificationResponse, error)
	DeleteNotification(id string) (*domain.NotificationResponse, error)
}
