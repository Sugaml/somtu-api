package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateBook creates a new Book
func (s *Service) CreateBook(ctx context.Context, req *domain.BookRequest) (*domain.BookResponse, error) {
	// Validate request
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Convert to Book model
	book := domain.Convert[domain.BookRequest, domain.Book](req)
	logrus.Info("Book to be created: ", book)
	// Create Book record
	result, err := s.repo.CreateBook(book)
	if err != nil {
		return nil, err
	}

	// Dynamically create BookCopy entries
	if req.AccessionType == "range" {
		for i := req.StartValue; i <= req.EndValue; i++ {
			copy := &domain.BookCopy{
				BookID:          result.ID,
				AccessionNumber: fmt.Sprintf("%03d", i),
				Status:          "available",
			}
			if _, err := s.repo.CreateBookCopy(copy); err != nil {
				// optionally log error but continue
				logrus.Error("Failed to create book copy: ", err)
			}
		}
	} else {
		copies := make([]*domain.BookCopy, result.TotalCopies)
		for i := uint(0); i < result.TotalCopies; i++ {
			copies[i] = &domain.BookCopy{
				BookID:          result.ID,
				AccessionNumber: fmt.Sprintf("%03d", i),
				Status:          "available",
			}
		}
		// Bulk insert copies
		for _, copy := range copies {
			if _, err := s.repo.CreateBookCopy(copy); err != nil {
				// optionally log error but continue
				logrus.Error("Failed to create book copy: ", err)
			}
		}
	}

	// Optional: Create notification and audit log
	userID, _ := getUserID(ctx)
	_, _ = s.repo.CreateNotification(&domain.Notification{
		Title:       fmt.Sprintf("Created new Book %s with %d copies.", result.Title, result.TotalCopies),
		Description: "create",
		UserID:      userID,
		Type:        "book",
		Action:      "create",
		Module:      "book",
		IsActive:    true,
	})

	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Created new Book %s with copies.", result.Title),
		Action:   "create",
		Data:     string(domain.ConvertToJson(result)),
		IsActive: true,
	})

	return domain.Convert[domain.Book, domain.BookResponse](result), nil
}

// ListBooks retrieves a list of Books
func (s *Service) ListBook(ctx context.Context, req *domain.BookListRequest) ([]*domain.BookResponse, int64, error) {
	var datas = []*domain.BookResponse{}
	results, count, err := s.repo.ListBook(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		data := domain.Convert[domain.Book, domain.BookResponse](result)
		data.AvailableCopies, _ = s.repo.GetAvailableCopies(data.ID)
		datas = append(datas, data)
	}
	return datas, count, nil
}

func (s *Service) GetBook(ctx context.Context, id string) (*domain.BookResponse, error) {
	result, err := s.repo.GetBook(id)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.Book, domain.BookResponse](result)
	return data, nil
}

func (s *Service) UpdateBook(ctx context.Context, id string, req *domain.BookUpdateRequest) (*domain.BookResponse, error) {
	if id == "" {
		return nil, errors.New("required Book id")
	}
	_, err := s.repo.GetBook(id)
	if err != nil {
		return nil, err
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	// update
	mp := req.NewUpdate()
	result, err := s.repo.UpdateBook(id, mp)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateNotification(&domain.Notification{
		Title:       fmt.Sprintf("Updated %s Book details.", result.Title),
		Description: "update",
		Type:        "book",
		Action:      "update",
		Module:      "book",
		UserID:      getUserID,
		IsActive:    true,
	})
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Updated %s Book details.", result.Title),
		Action:   "update",
		Data:     fmt.Sprint(req),
		UserID:   &getUserID,
		IsActive: true,
	})
	data := domain.Convert[domain.Book, domain.BookResponse](result)
	return data, nil
}

func (s *Service) DeleteBook(ctx context.Context, id string) (*domain.BookResponse, error) {
	result, err := s.repo.GetBook(id)
	if err != nil {
		return nil, err
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	copyExists, err := s.repo.IsBookCopiesByBookId(id)
	if err != nil {
		return nil, err
	}
	if copyExists {
		return nil, fmt.Errorf("book has copies cannot delete it")
	}
	CountBorrwedCopiesBookID, err := s.repo.CountBorrwedCopiesBookID(id)
	if err != nil {
		return nil, err
	}
	logrus.Info("CountBorrwedCopiesBookID :: ", CountBorrwedCopiesBookID)
	if CountBorrwedCopiesBookID > 0 {
		return nil, fmt.Errorf("book has %d copies borrowed cannot delete it", CountBorrwedCopiesBookID)
	}
	err = s.repo.DeleteBook(id)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Deleted %s parking area.", result.Title),
		Action:   "delete",
		UserID:   &getUserID,
		Data:     fmt.Sprint(result),
		IsActive: true,
	})
	return domain.Convert[domain.Book, domain.BookResponse](result), nil
}
