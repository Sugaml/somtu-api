package port

import (
	"github.com/sugaml/lms-api/internal/core/domain"
)

// type FineRepository interface is an interface for interacting with type Announcement-related data
type FineRepository interface {
	CreateFine(data *domain.Fine) (*domain.Fine, error)
	ListFine(req *domain.ListFineRequest) ([]*domain.Fine, int64, error)
	GetFine(id string) (*domain.Fine, error)
	UpdateFine(id string, req domain.Map) (*domain.Fine, error)
	DeleteFine(id string) error
}

// type FineService interface is an interface for interacting with type Announcement-related data
type FineService interface {
	CreateFine(data *domain.FineRequest) (*domain.FineResponse, error)
	ListFine(req *domain.ListFineRequest) ([]*domain.FineResponse, int64, error)
	GetFine(id string) (*domain.FineResponse, error)
	UpdateFine(id string, req *domain.UpdateFineRequest) (*domain.FineResponse, error)
	DeleteFine(id string) (*domain.FineResponse, error)
}
