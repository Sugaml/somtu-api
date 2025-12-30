package port

import (
	"context"

	"github.com/sugaml/lms-api/internal/core/domain"
)

//go:generate  mockgen -source=category.go -destination=../../adapter/storage/postgres/repository/mocks/mock_category_repository.go -package=mocks port.CategoryRepository

// CategoryRepository is an interface for interacting with category-related data
type ProgramRepository interface {
	CreateProgram(ctx context.Context, c *domain.Program) (*domain.Program, error)
	ListProgram(ctx context.Context, req *domain.ListProgramRequest) ([]*domain.Program, int64, error)
	GetbyNameProgram(ctx context.Context, name string) (*domain.Program, error)
	GetProgram(ctx context.Context, id string) (*domain.Program, error)
	UpdateProgram(ctx context.Context, id string, req domain.Map) error
	DeleteProgram(ctx context.Context, id string) error
}

// ProgramService is an interface for interacting with Program-related business logic
type ProgramService interface {
	//Creates a new Program
	CreateProgram(ctx context.Context, c *domain.ProgramRequest) (*domain.ProgramResponse, error)
	LisProgram(ctx context.Context, req *domain.ListProgramRequest) ([]domain.ProgramResponse, int64, error)
	GetProgram(ctx context.Context, id string) (*domain.ProgramResponse, error)
	UpdateProgram(ctx context.Context, id string, req *domain.ProgramUpdateRequest) (*domain.ProgramResponse, error)
	DeleteProgram(ctx context.Context, id string) error
}
