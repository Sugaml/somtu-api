package repository

import (
	"context"

	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateCategory creates a new Category record in the database
func (r *Repository) Create(ctx context.Context, data *domain.Category) (*domain.Category, error) {
	err := r.db.Model(&domain.Category{}).Create(&data).Take(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) GetbyName(ctx context.Context, name string) (*domain.Category, error) {
	data := &domain.Category{}
	err := r.db.
		Model(&domain.Category{}).
		Select("id, name, created_at, updated_at, weight, is_active").
		Where("name = ? AND is_active = true", name).
		Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) Get(ctx context.Context, id string) (*domain.Category, error) {
	data := &domain.Category{}
	err := r.db.
		Model(&domain.Category{}).
		Select("id, name, created_at, updated_at, weight, is_active").
		Where("id = ? AND is_active = true", id).
		Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 3. List All Categories
func (r *Repository) List(ctx context.Context, req *domain.ListCategoryRequest) ([]*domain.Category, int64, error) {
	var categories []*domain.Category
	count := 0
	f := r.db.Model(&domain.Category{})
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

// 4. Update Category
func (r *Repository) Update(ctx context.Context, id string, req domain.Map) error {
	category := &domain.Category{}
	err := r.db.Model(&domain.Category{}).Where("id = ?", id).Updates(category).Error
	if err != nil {
		return err
	}
	return nil
}

// 5. Delete Category
func (r *Repository) Delete(ctx context.Context, id string) error {
	return r.db.Model(&domain.Category{}).Where("id = ?", id).Delete(&domain.Category{}).Error
}
