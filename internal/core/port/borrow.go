package port

import (
	"context"

	"github.com/sugaml/lms-api/internal/core/domain"
)

// type BorrowRepository interface is an interface for interacting with type Announcement-related data
type BorrowRepository interface {
	CreateBorrow(data *domain.BorrowedBook) (*domain.BorrowedBook, error)
	ListBorrow(req *domain.ListBorrowedBookRequest) ([]*domain.BorrowedBook, int64, error)
	GetBorrow(id string) (*domain.BorrowedBook, error)
	GetAvailableCopies(bookID string) (uint, error)
	GetBookBorrowByUserID(user_id string) ([]*domain.BorrowedBook, error)
	IsBookBorrowByUserID(user_id string, book_id string) bool
	CountAllBookBorrwedCopies() (int64, error)
	CountBorrwedCopiesBookID(bookID string) (int64, error)
	CountBorrwedCopiesUserID(userID string) (int64, error)
	UpdateBorrow(id string, req domain.Map) (*domain.BorrowedBook, error)
	DeleteBorrow(id string) error
}

// type BorrowService interface is an interface for interacting with type Announcement-related data
type BorrowService interface {
	CreateBorrow(ctx context.Context, data *domain.BorrowedBookRequest) (*domain.BorrowedBookResponse, error)
	ListBorrow(ctx context.Context, req *domain.ListBorrowedBookRequest) ([]*domain.BorrowedBookResponse, int64, error)
	GetStudentsBorrowBook(ctx context.Context, id string) ([]*domain.BorrowedBookResponse, error)
	GetBorrow(ctx context.Context, id string) (*domain.BorrowedBookResponse, error)
	UpdateBorrow(ctx context.Context, id string, req *domain.UpdateBorrowedBookRequest) (*domain.BorrowedBookResponse, error)
	DeleteBorrow(ctx context.Context, id string) (*domain.BorrowedBookResponse, error)
}
