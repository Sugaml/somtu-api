package domain

import "time"

type Fine struct {
	BaseModel
	UserID         int        `gorm:"not null" json:"user_id"`
	BorrowedBookID int        `gorm:"column:borrowed_book_id;not null" json:"borrowed_book_id"`
	Amount         int        `gorm:"not null" json:"amount"` // in paisa
	Reason         string     `gorm:"not null" json:"reason"`
	Status         string     `gorm:"not null" json:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `gorm:"column:paid_at" json:"paid_at"`
	IsActive       bool       `gorm:"column:is_active;default:false" json:"is_active"`
}

type FineRequest struct {
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UserID         int        `json:"user_id"`
	BorrowedBookID int        `json:"borrowed_book_id"`
	Amount         int        `json:"amount"` // in paisa
	Reason         string     `json:"reason"`
	Status         string     `json:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `json:"paid_at"`
}

type UpdateFineRequest struct {
	UserID         int        `json:"user_id"`
	BorrowedBookID int        `json:"borrowed_book_id"`
	Amount         int        `json:"amount"` // in paisa
	Reason         string     `json:"reason"`
	Status         string     `json:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `json:"paid_at"`
}

type ListFineRequest struct {
	ListRequest
	UserID         int        `form:"user_id"`
	BorrowedBookID int        `form:"borrowed_book_id"`
	Amount         int        `form:"amount"` // in paisa
	Reason         string     `form:"reason"`
	Status         string     `form:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `form:"paid_at"`
}

type FineResponse struct {
	ID             string     `json:"id"`
	CreatedAt      time.Time  `json:"created_at"`
	UserID         int        `json:"user_id"`
	BorrowedBookID int        `json:"borrowed_book_id"`
	Amount         int        `json:"amount"` // in paisa
	Reason         string     `json:"reason"`
	Status         string     `json:"status"` // 'pending' | 'paid'
	PaidAt         *time.Time `json:"paid_at"`
}

func (u *FineRequest) Validate() error {
	return nil
}

func (r *UpdateFineRequest) NewUpdate() Map {
	return nil
}
