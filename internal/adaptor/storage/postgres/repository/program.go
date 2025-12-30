package repository

import (
	"context"

	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateProgram creates a new Program record in the database
func (r *Repository) CreateProgram(ctx context.Context, data *domain.Program) (*domain.Program, error) {
	err := r.db.Model(&domain.Program{}).Create(&data).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) GetbyNameProgram(ctx context.Context, name string) (*domain.Program, error) {
	data := &domain.Program{}
	err := r.db.
		Model(&domain.Program{}).
		Select("id, name, created_at, updated_at, weight, is_active").
		Where("name = ? AND is_active = true", name).
		Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) GetProgram(ctx context.Context, id string) (*domain.Program, error) {
	data := &domain.Program{}
	err := r.db.
		Model(&domain.Program{}).
		Select("id, name, created_at, updated_at, weight, is_active").
		Where("id = ? AND is_active = true", id).
		Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 3. List All Categories
func (r *Repository) ListProgram(ctx context.Context, req *domain.ListProgramRequest) ([]*domain.Program, int64, error) {
	var categories []*domain.Program
	count := 0
	f := r.db.Model(&domain.Program{})
	f = f.Where("lower(name) LIKE lower(?)", "%"+req.Query+"%")
	err := f.Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&categories).Error
	if err != nil {
		return nil, int64(count), err
	}
	return categories, int64(count), nil
}

// 4. Update Program
func (r *Repository) UpdateProgram(ctx context.Context, id string, req domain.Map) error {
	Program := &domain.Program{}
	err := r.db.Model(&domain.Program{}).Where("id = ?", id).Updates(Program).Error
	if err != nil {
		return err
	}
	return nil
}

// 5. Delete Program
func (r *Repository) DeleteProgram(ctx context.Context, id string) error {
	return r.db.Model(&domain.Program{}).Where("id = ?", id).Delete(&domain.Program{}).Error
}
