package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateBorrowBook creates a new BorrowedBook
func (s *Service) CreateBorrow(ctx context.Context, req *domain.BorrowedBookRequest) (*domain.BorrowedBookResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	isBookBorrowd := s.repo.IsBookBorrowByUserID(req.UserID, req.BookCopyID)
	if isBookBorrowd {
		return nil, errors.New("book already borrowed")
	}
	bookCopy, err := s.repo.GetBookCopy(req.BookCopyID)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetUser(req.UserID)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.BorrowedBookRequest, domain.BorrowedBook](req)
	if data.Status == "" {
		data.Status = "pending"
	}
	data.BookID = bookCopy.BookID
	data.IsActive = true
	result, err := s.repo.CreateBorrow(data)
	if err != nil {
		return nil, err
	}
	if data.Status == "borrowed" {
		data.Status = "issued"
		_, err = s.repo.UpdateBookCopy(bookCopy.ID, domain.Map{"status": "borrowed"})
		if err != nil {
			return nil, err
		}
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("Book %s Accession Number %s has been issued to %s", bookCopy.Book.Title, bookCopy.AccessionNumber, user.FullName),
			UserID:   &getUserID,
			Action:   "issue",
			Data:     fmt.Sprint(req),
			IsActive: true,
		})
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("%s book has %s to %s", bookCopy.Book.Title, result.Status, user.FullName),
			UserID:   &getUserID,
			Action:   "issue",
			Data:     fmt.Sprint(req),
			IsActive: true,
		})
	}
	if data.Status == "pending" {
		data.Status = "requested"
		s.repo.CreateNotification(&domain.Notification{
			Title:    fmt.Sprintf("%s book has %s by %s", bookCopy.Book.Title, data.Status, user.FullName),
			UserID:   getUserID,
			Module:   "borrow",
			Action:   "borrow",
			IsActive: true,
		})
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("%s book has %s by %s", bookCopy.Book.Title, result.Status, user.FullName),
			UserID:   &getUserID,
			Action:   "create",
			Data:     fmt.Sprint(req),
			IsActive: true,
		})
	}
	return domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result), nil
}

// ListBorrowedBooks retrieves a list of BorrowedBooks
func (s *Service) ListBorrow(ctx context.Context, req *domain.ListBorrowedBookRequest) ([]*domain.BorrowedBookResponse, int64, error) {
	var datas = []*domain.BorrowedBookResponse{}
	results, count, err := s.repo.ListBorrow(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		data := domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result)
		datas = append(datas, data)
	}
	return datas, count, nil
}

func (s *Service) GetBorrow(ctx context.Context, id string) (*domain.BorrowedBookResponse, error) {
	result, err := s.repo.GetBorrow(id)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result)
	return data, nil
}

func (s *Service) GetStudentsBorrowBook(ctx context.Context, id string) ([]*domain.BorrowedBookResponse, error) {
	var datas = []*domain.BorrowedBookResponse{}
	results, err := s.repo.GetBookBorrowByUserID(id)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		data := domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result)
		datas = append(datas, data)
	}
	return datas, nil
}

func (s *Service) UpdateBorrow(ctx context.Context, id string, req *domain.UpdateBorrowedBookRequest) (*domain.BorrowedBookResponse, error) {
	if id == "" {
		return nil, errors.New("required borrow id")
	}
	borrow, err := s.repo.GetBorrow(id)
	if err != nil {
		return nil, err
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	bookCopy, err := s.repo.GetBookCopy(borrow.BookCopyID)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.GetUser(borrow.UserID)
	if err != nil {
		return nil, err
	}
	// update
	mp := req.NewUpdate()
	logrus.Info(mp)
	result, err := s.repo.UpdateBorrow(id, mp)
	if err != nil {
		return nil, err
	}
	if result.Status == "borrowed" {
		result.Status = "issued"
		_, err = s.repo.UpdateBookCopy(bookCopy.ID, domain.Map{"status": "borrowed"})
		if err != nil {
			return nil, err
		}
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("Book %s Accession Number %s has been issued to %s", bookCopy.Book.Title, bookCopy.AccessionNumber, user.FullName),
			UserID:   &getUserID,
			Action:   "issue",
			Data:     fmt.Sprint(req),
			IsActive: true,
		})
		_, _ = s.repo.CreateNotification(&domain.Notification{
			Title:    fmt.Sprintf("%s book has %s to %s", bookCopy.Book.Title, result.Status, user.FullName),
			UserID:   user.ID,
			Module:   "borrow",
			Action:   "issue",
			IsActive: true,
		})
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("%s book has %s to %s", bookCopy.Book.Title, result.Status, user.FullName),
			UserID:   &getUserID,
			Action:   "issue",
			Data:     fmt.Sprint(req),
			IsActive: true,
		})
	}
	if result.Status == "returned" {
		result.Status = "returned"
		_, err = s.repo.UpdateBookCopy(bookCopy.ID, domain.Map{"status": "available"})
		if err != nil {
			return nil, err
		}
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("Book %s Accession Number %s has been available", bookCopy.Book.Title, bookCopy.AccessionNumber),
			UserID:   &getUserID,
			Action:   "issue",
			Data:     fmt.Sprint(req),
			IsActive: true,
		})
		_, _ = s.repo.CreateNotification(&domain.Notification{
			Title:    fmt.Sprintf("%s book has %s by %s", bookCopy.Book.Title, result.Status, user.FullName),
			UserID:   user.ID,
			Module:   "borrow",
			Action:   "return",
			IsActive: true,
		})
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("%s book has %s by %s", bookCopy.Book.Title, result.Status, user.FullName),
			UserID:   &getUserID,
			Action:   "update",
			Data:     fmt.Sprint(req),
			IsActive: true,
		})
	}
	data := domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result)
	return data, nil
}

func (s *Service) DeleteBorrow(ctx context.Context, id string) (*domain.BorrowedBookResponse, error) {
	result, err := s.repo.GetBorrow(id)
	if err != nil {
		return nil, err
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.repo.DeleteBorrow(id)
	if err != nil {
		return nil, err
	}
	s.repo.CreateNotification(&domain.Notification{
		Title:    fmt.Sprintf("%s book has been deleted by %s", result.BookCopy.Book.Title, result.Student.FullName),
		UserID:   result.UserID,
		Module:   "borrow",
		Action:   "delete",
		IsActive: true,
	})
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("%s book has been deleted by %s", result.BookCopy.Book.Title, result.Student.FullName),
		UserID:   &getUserID,
		Action:   "delete",
		Data:     fmt.Sprint(result),
		IsActive: true,
	})
	return domain.Convert[domain.BorrowedBook, domain.BorrowedBookResponse](result), nil
}
