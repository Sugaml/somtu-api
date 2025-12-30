package repository

import (
	"github.com/sugaml/lms-api/internal/core/port"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// NewAnnouncementRepository creates a new Announcement repository instance
func NewRepository(db *gorm.DB) port.Repository {
	return &Repository{
		db,
	}
}
