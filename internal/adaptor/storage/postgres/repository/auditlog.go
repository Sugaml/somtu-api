package repository

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) CreateAuditLog(data *domain.AuditLog) (*domain.AuditLog, error) {
	if err := r.db.Model(&domain.AuditLog{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) ListAuditLog(req *domain.ListAuditLogRequest) ([]*domain.AuditLog, int64, error) {
	var datas []*domain.AuditLog
	var count int64
	f := r.db.Model(&domain.AuditLog{})
	if req.Query != "" {
		f = f.Where("lower(title) LIKE lower(?)", "%"+req.Query+"%")
	}
	if req.UserID != "" {
		f = f.Where("user_id = ?", req.UserID)
	}
	if req.Action != "" {
		f = f.Where("action = ?", req.Action)
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

func (r *Repository) GetAuditLog(id string) (*domain.AuditLog, error) {
	var data domain.AuditLog
	if err := r.db.Model(&domain.AuditLog{}).
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) UpdateAuditLog(id string, req domain.Map) (*domain.AuditLog, error) {
	if id == "" {
		return nil, errors.New("required AuditLog id")
	}
	data := &domain.AuditLog{}
	err := r.db.Model(&domain.AuditLog{}).Where("id = ?", id).Updates(req.ToMap()).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteAuditLog(id string) error {
	return r.db.Model(&domain.AuditLog{}).Where("id = ?", id).Delete(&domain.AuditLog{}).Error
}
