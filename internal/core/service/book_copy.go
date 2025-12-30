package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (s *Service) CreateBookCopy(ctx context.Context, req *domain.AddBookCopiesRequest) (*domain.BookCopyResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	book, err := s.repo.GetBook(req.BookID)
	if err != nil {
		return nil, errors.New("book not found")
	}
	total := book.TotalCopies + req.AddCopies
	_, err = s.repo.UpdateBook(req.BookID, domain.Map{"total_copies": total})
	if err != nil {
		return nil, err
	}
	var result *domain.BookCopy
	for i := req.StartAccessionNumber; i < req.StartAccessionNumber+req.AddCopies; i++ {
		data := &domain.BookCopy{
			BookID:          req.BookID,
			AccessionNumber: fmt.Sprintf("%03d", i),
			Status:          "available",
		}
		result, err = s.repo.CreateBookCopy(data)
		if err != nil {
			return nil, err
		}
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("Created new copy %s of BookID %s", result.AccessionNumber, result.BookID),
			Action:   "create",
			Data:     fmt.Sprint(result),
			IsActive: true,
		})
		userID, _ := getUserID(ctx)
		_, _ = s.repo.CreateNotification(&domain.Notification{
			Title:       fmt.Sprintf("Created new copy %s of BookID %s", result.AccessionNumber, result.BookID),
			Description: "create",
			UserID:      userID,
			Type:        "book_copy",
			Action:      "create",
			Module:      "book_copy",
			IsActive:    true,
		})
	}
	return domain.Convert[domain.BookCopy, domain.BookCopyResponse](result), nil
}

func (s *Service) ListBookCopies(ctx context.Context, req *domain.BookCopyListRequest) ([]*domain.BookCopyResponse, int64, error) {
	var datas []*domain.BookCopyResponse
	results, count, err := s.repo.ListBookCopies(req)
	if err != nil {
		return nil, count, err
	}

	for _, r := range results {
		data := domain.Convert[domain.BookCopy, domain.BookCopyResponse](r)
		datas = append(datas, data)
	}

	return datas, count, nil
}

func (s *Service) ListBookCopiesByBookId(ctx context.Context, bookId string, req *domain.BookCopyListRequest) ([]*domain.BookCopyResponse, int64, error) {
	var datas []*domain.BookCopyResponse
	results, count, err := s.repo.ListBookCopiesByBookId(bookId, req)
	if err != nil {
		return nil, count, err
	}

	for _, r := range results {
		data := domain.Convert[domain.BookCopy, domain.BookCopyResponse](r)
		datas = append(datas, data)
	}

	return datas, count, nil
}

func (s *Service) GetBookCopy(ctx context.Context, id string) (*domain.BookCopyResponse, error) {
	result, err := s.repo.GetBookCopy(id)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.BookCopy, domain.BookCopyResponse](result)
	return data, nil
}

func (s *Service) UpdateBookCopy(ctx context.Context, id string, req *domain.BookCopyUpdateRequest) (*domain.BookCopyResponse, error) {
	if id == "" {
		return nil, errors.New("required BookCopy id")
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.repo.GetBookCopy(id)
	if err != nil {
		return nil, err
	}

	mp := req.NewUpdate()
	result, err := s.repo.UpdateBookCopy(id, mp)
	if err != nil {
		return nil, err
	}

	_, _ = s.repo.CreateNotification(&domain.Notification{
		Title:       fmt.Sprintf("Updated copy %s of BookID %s", result.AccessionNumber, result.BookID),
		Description: "update",
		UserID:      getUserID,
		Type:        "book_copy",
		Action:      "update",
		Module:      "book_copy",
		IsActive:    true,
	})
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Updated copy %s of BookID %s", result.AccessionNumber, result.BookID),
		Action:   "update",
		UserID:   &getUserID,
		Data:     fmt.Sprint(req),
		IsActive: true,
	})

	return domain.Convert[domain.BookCopy, domain.BookCopyResponse](result), nil
}

func (s *Service) DeleteBookCopy(ctx context.Context, id string) (*domain.BookCopyResponse, error) {
	result, err := s.repo.GetBookCopy(id)
	if err != nil {
		return nil, err
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	// Optional: check if this copy is currently issued
	borrowedCount, err := s.repo.CountBorrowedCopyID(id)
	if err != nil {
		return nil, err
	}
	if borrowedCount > 0 {
		return nil, fmt.Errorf("this copy is currently borrowed, cannot delete")
	}

	err = s.repo.DeleteBookCopy(id)
	if err != nil {
		return nil, err
	}

	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Deleted copy %s of BookID %s", result.AccessionNumber, result.BookID),
		Action:   "delete",
		UserID:   &getUserID,
		Data:     fmt.Sprint(result),
		IsActive: true,
	})

	return domain.Convert[domain.BookCopy, domain.BookCopyResponse](result), nil
}
