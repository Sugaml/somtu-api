package repository

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) CreateFine(data *domain.Fine) (*domain.Fine, error) {
	if err := r.db.Model(&domain.Fine{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) ListFine(req *domain.ListFineRequest) ([]*domain.Fine, int64, error) {
	var datas []*domain.Fine
	var count int64
	f := r.db.Model(&domain.Fine{})
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

func (r *Repository) GetFine(id string) (*domain.Fine, error) {
	var data domain.Fine
	if err := r.db.Model(&domain.Fine{}).
		Preload("Fine").
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) UpdateFine(id string, req domain.Map) (*domain.Fine, error) {
	if id == "" {
		return nil, errors.New("required Fine id")
	}
	data := &domain.Fine{}
	err := r.db.Model(&domain.Fine{}).Where("id = ?", id).Updates(req.ToMap()).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteFine(id string) error {
	return r.db.Model(&domain.Fine{}).Where("id = ?", id).Delete(&domain.Fine{}).Error
}
