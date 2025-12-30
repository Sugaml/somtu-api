package repository

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) CreateNotification(data *domain.Notification) (*domain.Notification, error) {
	if err := r.db.Model(&domain.Notification{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) ListNotification(req *domain.ListNotificationRequest) ([]*domain.Notification, int64, error) {
	var datas []*domain.Notification
	var count int64
	f := r.db.Model(&domain.Notification{})
	if req.Query != "" {
		req.SortColumn = "score desc, " + req.SortColumn
	}
	err := f.Count(&count).
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, nil
}

func (r *Repository) GetNotification(id string) (*domain.Notification, error) {
	var data domain.Notification
	if err := r.db.Model(&domain.Notification{}).
		Preload("Notification").
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) UpdateNotification(id string, req domain.Map) (*domain.Notification, error) {
	if id == "" {
		return nil, errors.New("required notification id")
	}
	data := &domain.Notification{}
	err := r.db.Model(&domain.Notification{}).Where("id = ?", id).Updates(req.ToMap()).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) ReadAllNotification(req domain.Map) (*domain.Notification, error) {
	data := &domain.Notification{}
	err := r.db.Model(&domain.Notification{}).Where("is_read = ?", false).Updates(req.ToMap()).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteNotification(id string) error {
	return r.db.Model(&domain.Notification{}).Where("id = ?", id).Delete(&domain.Notification{}).Error
}
