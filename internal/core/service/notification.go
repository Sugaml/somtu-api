package service

import (
	"errors"
	"fmt"

	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateNotification creates a new Notification
func (s *Service) CreateNotification(req *domain.NotificationRequest) (*domain.NotificationResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.NotificationRequest, domain.Notification](req)
	result, err := s.repo.CreateNotification(data)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Created new Notification %s.", result.Title),
		Action:   "create",
		Data:     string(domain.ConvertToJson(result)),
		IsActive: true,
	})
	return domain.Convert[domain.Notification, domain.NotificationResponse](result), nil
}

// ListNotifications retrieves a list of Notifications
func (s *Service) ListNotification(req *domain.ListNotificationRequest) ([]*domain.NotificationResponse, int64, error) {
	var datas = []*domain.NotificationResponse{}
	results, count, err := s.repo.ListNotification(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		data := domain.Convert[domain.Notification, domain.NotificationResponse](result)
		datas = append(datas, data)
	}
	return datas, count, nil
}

func (s *Service) GetNotification(id string) (*domain.NotificationResponse, error) {
	result, err := s.repo.GetNotification(id)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.Notification, domain.NotificationResponse](result)
	return data, nil
}

func (s *Service) UpdateNotification(id string, req *domain.UpdateNotificationRequest) (*domain.NotificationResponse, error) {
	if id == "" {
		return nil, errors.New("required Notification id")
	}
	_, err := s.repo.GetNotification(id)
	if err != nil {
		return nil, err
	}

	// update
	mp := req.NewUpdate()
	result, err := s.repo.UpdateNotification(id, mp)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Updated %s Notification details.", result.Title),
		Action:   "update",
		Data:     fmt.Sprint(req),
		IsActive: true,
	})
	data := domain.Convert[domain.Notification, domain.NotificationResponse](result)
	return data, nil
}

func (s *Service) ReadAllNotification() (*domain.NotificationResponse, error) {
	mp := map[string]interface{}{"is_read": true}
	result, err := s.repo.ReadAllNotification(mp)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    "Read all Notifications.",
		Action:   "read_all",
		Data:     fmt.Sprint(result),
		IsActive: true,
	})
	if result == nil {
		return nil, errors.New("no unread notifications found")
	}
	data := domain.Convert[domain.Notification, domain.NotificationResponse](result)
	return data, nil
}

func (s *Service) DeleteNotification(id string) (*domain.NotificationResponse, error) {
	result, err := s.repo.GetNotification(id)
	if err != nil {
		return nil, err
	}
	err = s.repo.DeleteNotification(id)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Deleted %s parking area.", result.Title),
		Action:   "delete",
		Data:     fmt.Sprint(result),
		IsActive: true,
	})
	return domain.Convert[domain.Notification, domain.NotificationResponse](result), nil
}
