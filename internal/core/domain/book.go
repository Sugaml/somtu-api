package domain

import (
	"errors"
	"time"
)

type Book struct {
	BaseModel
	Title  string `gorm:"type:varchar(255);not null" json:"title"`
	Author string `gorm:"type:varchar(255);not null" json:"author"`
	// ISBN        string    `gorm:"type:varchar(20);null" json:"isbn"`
	Publisher   string    `gorm:"type:varchar(255);null" json:"publisher"`
	Edition     string    `gorm:"type:varchar(100)" json:"edition,omitempty"`
	CategoryID  string    `gorm:"not null" json:"category_id"`
	Description string    `gorm:"type:text" json:"description"`
	CoverImage  string    `gorm:"type:text" json:"cover_image,omitempty"`
	TotalPages  uint      `json:"total_pages"`
	TotalCopies uint      `gorm:"not null;default:1" json:"total_copies"`
	IsActive    bool      `gorm:"default:true" json:"is_active"`
	Keywords    string    `json:"keywords"`
	Tags        string    `json:"tags"`
	Category    *Category `gorm:"foreignkey:ID;references:CategoryID" json:"category,omitempty"`
	// Relations
	Copies []BookCopy `gorm:"foreignKey:BookID" json:"copies,omitempty"`
}

type BookCopy struct {
	BaseModel
	BookID          string `gorm:"not null" json:"book_id"` // FK to Book
	AccessionNumber string `gorm:"type:varchar(50);unique;not null" json:"accession_number"`
	Status          string `gorm:"type:varchar(20);not null;default:'available'" json:"status"`
	Book            *Book  `gorm:"foreignkey:ID;references:BookID" json:"book,omitempty"`
	Labels          string `json:"labels"`
	Remarks         string `json:"remarks"`
	// Relations
	BorrowedBooks []BorrowedBook `gorm:"foreignKey:BookCopyID" json:"borrowed_books,omitempty"`
}

type AddBookCopiesRequest struct {
	BookID               string `json:"book_id"`
	Remarks              string `json:"remarks"`
	AddCopies            uint   `json:"add_copies"`
	StartAccessionNumber uint   `json:"start_accession_number"`
	EndAccessionNumber   uint   `json:"end_accession_number"`
}

type BookCopyRequest struct {
	Remarks         string `json:"remarks"`
	BookID          string `gorm:"not null" json:"book_id"` // FK to Book
	AccessionNumber string `gorm:"type:varchar(50);unique;not null" json:"accession_number"`
	Status          string `gorm:"type:varchar(20);not null;default:'available'" json:"status"`
}

type BookCopyUpdateRequest struct {
	BookID          string `gorm:"not null" json:"book_id"` // FK to Book
	AccessionNumber string `gorm:"type:varchar(50);unique;not null" json:"accession_number"`
	Remarks         string `json:"remarks"`
	Status          string `gorm:"type:varchar(20);not null;default:'available'" json:"status"`
}

type BookCopyResponse struct {
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"created_at"`
	BookID          string        `json:"book_id"` // FK to Book
	AccessionNumber string        `json:"accession_number"`
	Status          string        `json:"status"`
	Remarks         string        `json:"remarks"`
	Book            *BookResponse `json:"book,omitempty"`
	// Relations
	BorrowedBooks []BorrowedBookResponse `json:"borrowed_books,omitempty"`
}

type BookCopyListRequest struct {
	ListRequest
	BookID          string `form:"book_id"`
	AccessionNumber string `form:"accession_number"`
	Status          string `form:"status"` // available, issued, reserved
}

type BookRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	ISBN          string `json:"isbn"`
	Keywords      string `json:"keywords"`
	Tags          string `json:"tags"`
	Publisher     string `json:"publisher"`
	AccessionType string `json:"accession_type"`
	StartValue    int    `json:"start_value"`
	EndValue      int    `json:"end_value"`
	Edition       string `json:"edition,omitempty"`
	CategoryID    string `json:"category_id"`
	TotalCopies   uint   `json:"total_copies"`
	TotalPages    uint   `json:"total_pages"`
	Description   string `json:"description"`
	CoverImage    string `json:"cover_image"`
}

type BookListRequest struct {
	ListRequest
	Title       string `form:"title"`
	Author      string `form:"author"`
	ISBN        string `form:"isbn"`
	Publisher   string `form:"publisher"`
	Edition     string `form:"edition,omitempty"`
	Category    string `form:"category"`
	Program     string `form:"program"`
	TotalCopies uint   `form:"total_copies"`
	Description string `form:"description"`
	CoverImage  string `form:"cover_image"`
}

type BookUpdateRequest struct {
	Title       *string `json:"title"`
	Author      *string `json:"author"`
	ISBN        *string `json:"isbn"`
	Publisher   *string `json:"publisher"`
	Keywords    string  `json:"keywords"`
	Tags        string  `json:"tags"`
	Edition     *string `json:"edition,omitempty"`
	Category    *string `json:"category"`
	Program     *string `json:"program"`
	TotalCopies *uint   `json:"total_copies"`
	TotalPages  *uint   `json:"total_pages"`
	Description *string `json:"description"`
	CoverImage  *string `json:"cover_image"`
}

type BookAllUpdateRequest struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	Publisher       string `json:"publisher"`
	Edition         string `json:"edition"`
	Category        string `json:"category"`
	Program         string `json:"program"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
	Description     string `json:"description"`
	CoverImage      string `json:"cover_image"`
}

type BookResponse struct {
	ID              string             `json:"id"`
	CreatedAt       time.Time          `json:"created_at"`
	Title           string             `json:"title"`
	Author          string             `json:"author"`
	ISBN            string             `json:"isbn"`
	Publisher       string             `json:"publisher"`
	Edition         string             `json:"edition"`
	Keywords        string             `json:"keywords"`
	Tags            string             `json:"tags"`
	ProgramID       string             `json:"program_id"`
	CategoryID      string             `json:"category_id"`
	Program         ProgramResponse    `json:"program"`
	TotalCopies     uint               `json:"total_copies"`
	AvailableCopies uint               `json:"available_copies"`
	Description     string             `json:"description"`
	CoverImage      string             `json:"cover_image"`
	Category        *CategoryResponse  `json:"category,omitempty"`
	Programs        *ProgramResponse   `json:"programs,omitempty"`
	Copies          []BookCopyResponse `json:"copies,omitempty"`
}

func (r *BookRequest) Validate() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.Author == "" {
		return errors.New("author is required")
	}
	return nil
}

func (r *BookUpdateRequest) NewUpdate() Map {
	mp := map[string]interface{}{}
	if r.Title != nil {
		mp["title"] = *r.Title
	}
	if r.Author != nil {
		mp["author"] = *r.Author
	}
	if r.ISBN != nil {
		mp["isbn"] = *r.ISBN
	}
	if r.Publisher != nil {
		mp["publisher"] = *r.Publisher
	}
	if r.Keywords != "" {
		mp["keywords"] = r.Keywords
	}
	if r.Tags != "" {
		mp["tags"] = r.Tags
	}
	if r.Edition != nil {
		mp["edition"] = *r.Edition
	}
	if r.Description != nil {
		mp["description"] = *r.Description
	}
	if r.CoverImage != nil {
		mp["cover_image"] = *r.CoverImage
	}
	if r.Category != nil {
		mp["category"] = *r.Category
	}
	if r.Program != nil {
		mp["program"] = *r.Program
	}
	if r.TotalCopies != nil {
		mp["total_copies"] = *r.TotalCopies
	}
	if r.TotalPages != nil {
		mp["total_pages"] = *r.TotalPages
	}
	return mp
}

func (r *AddBookCopiesRequest) Validate() error {
	if r.BookID == "" {
		return errors.New("book_id is required")
	}
	// if r.StartAccessionNumber == 0 {
	// 	return errors.New("start_accession_number is required")
	// }
	// if r.EndAccessionNumber == 0 {
	// 	return errors.New("end_accession_number is required")
	// }
	return nil
}

func (r *BookCopyUpdateRequest) NewUpdate() Map {
	mp := map[string]interface{}{}
	if r.BookID != "" {
		mp["book_id"] = r.BookID
	}
	if r.AccessionNumber != "" {
		mp["accession_number"] = r.AccessionNumber
	}
	if r.Status != "" {
		mp["status"] = r.Status
	}
	return mp
}
