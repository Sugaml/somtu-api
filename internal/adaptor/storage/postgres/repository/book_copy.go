package repository

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) CreateBookCopy(copy *domain.BookCopy) (*domain.BookCopy, error) {
	if err := r.db.Model(&domain.BookCopy{}).Create(&copy).Error; err != nil {
		return nil, err
	}
	return copy, nil
}

func (r *Repository) ListBookCopies(req *domain.BookCopyListRequest) ([]*domain.BookCopy, int64, error) {
	var copies []*domain.BookCopy
	var count int64
	f := r.db.Model(&domain.BookCopy{})

	if req.Query != "" {
		req.SortColumn = "created_at desc, " + req.SortColumn
	}
	if req.BookID != "" {
		f = f.Where("book_id = ?", req.BookID)
	}
	if req.AccessionNumber != "" {
		f = f.Where("accession_number = ?", req.AccessionNumber)
	}
	if req.Status != "" {
		f = f.Where("status = ?", req.Status)
	}
	err := f.Count(&count).Preload("Book").
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&copies).Error
	if err != nil {
		return nil, count, err
	}
	return copies, count, nil
}

func (r *Repository) IsBookCopiesByBookId(bookId string) (bool, error) {
	var count int64
	if err := r.db.Model(&domain.BookCopy{}).Where("book_id = ?", bookId).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *Repository) ListBookCopiesByBookId(bookId string, req *domain.BookCopyListRequest) ([]*domain.BookCopy, int64, error) {
	var copies []*domain.BookCopy
	var count int64
	f := r.db.Model(&domain.BookCopy{})
	if req.Query != "" {
		req.SortColumn = "created_at desc, " + req.SortColumn
	}
	if req.Status != "" {
		f = f.Where("status = ?", req.Status)
	}
	err := f.Where("book_id = ?", bookId).Count(&count).Preload("Book").
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&copies).Error
	if err != nil {
		return nil, count, err
	}
	return copies, count, nil
}

func (r *Repository) GetBookCopy(id string) (*domain.BookCopy, error) {
	var copy domain.BookCopy
	if err := r.db.Model(&domain.BookCopy{}).Preload("Book").
		Take(&copy, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &copy, nil
}

func (r *Repository) UpdateBookCopy(id string, req domain.Map) (*domain.BookCopy, error) {
	if id == "" {
		return nil, errors.New("required book copy id")
	}
	copy := &domain.BookCopy{}
	err := r.db.Model(&domain.BookCopy{}).
		Where("id = ?", id).
		Updates(req.ToMap()).
		Take(&copy).Error
	if err != nil {
		return nil, err
	}
	return copy, nil
}

func (r *Repository) DeleteBookCopy(id string) error {
	return r.db.Model(&domain.BookCopy{}).Where("id = ?", id).Delete(&domain.BookCopy{}).Error
}

func (r *Repository) CountBorrowedCopyID(bookCopyID string) (int64, error) {
	var count int64
	err := r.db.Model(&domain.BorrowedBook{}).
		Where("book_copy_id = ? AND status IN ?", bookCopyID, []string{"borrowed", "pending", "overdue"}).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
