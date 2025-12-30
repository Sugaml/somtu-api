package port

import (
	"context"

	"github.com/sugaml/lms-api/internal/core/domain"
)

//go:generate  mockgen -source=category.go -destination=../../adapter/storage/postgres/repository/mocks/mock_category_repository.go -package=mocks port.CategoryRepository

// CategoryRepository is an interface for interacting with category-related data
type CategoryRepository interface {
	Create(ctx context.Context, c *domain.Category) (*domain.Category, error)
	List(ctx context.Context, req *domain.ListCategoryRequest) ([]*domain.Category, int64, error)
	GetbyName(ctx context.Context, name string) (*domain.Category, error)
	Get(ctx context.Context, id string) (*domain.Category, error)
	Update(ctx context.Context, id string, req domain.Map) error
	Delete(ctx context.Context, id string) error
}

// CategoryService is an interface for interacting with category-related business logic
type CategoryService interface {
	//Creates a new category
	Create(ctx context.Context, c *domain.CategoryRequest) (*domain.CategoryResponse, error)
	List(ctx context.Context, req *domain.ListCategoryRequest) ([]domain.CategoryResponse, int64, error)
	Get(ctx context.Context, id string) (*domain.CategoryResponse, error)
	Update(ctx context.Context, id string, req *domain.CategoryUpdateRequest) (*domain.CategoryResponse, error)
	Delete(ctx context.Context, id string) error
}
