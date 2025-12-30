package repository

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) CreateBook(data *domain.Book) (*domain.Book, error) {
	if err := r.db.Model(&domain.Book{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) ListBook(req *domain.BookListRequest) ([]*domain.Book, int64, error) {
	var datas []*domain.Book
	var count int64
	f := r.db.Model(&domain.Book{})
	if req.Query != "" {
		req.SortColumn = "score desc, " + req.SortColumn
	}
	if req.Title != "" {
		f = f.Where("title ILIKE ?", "%"+req.Title+"%") // Use ILIKE for case-insensitive search (PostgreSQL)
	}
	err := f.Count(&count).Preload("Category").
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, nil
}

func (r *Repository) GetBook(id string) (*domain.Book, error) {
	var data domain.Book
	if err := r.db.Model(&domain.Book{}).
		Preload("Copies").
		Preload("Category").
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) UpdateBook(id string, req domain.Map) (*domain.Book, error) {
	if id == "" {
		return nil, errors.New("required book id")
	}
	data := &domain.Book{}
	err := r.db.Model(&domain.Book{}).Where("id = ?", id).Updates(req.ToMap()).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteBook(id string) error {
	return r.db.Model(&domain.Book{}).Where("id = ?", id).Delete(&domain.Book{}).Error
}
