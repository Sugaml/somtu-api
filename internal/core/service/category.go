package service

import (
	"context"
	"fmt"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (s *Service) Create(ctx context.Context, req *domain.CategoryRequest) (*domain.CategoryResponse, error) {
	category := &domain.Category{}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	category.NewCategory(req)
	category, err = s.repo.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Created %s Category", category.Name),
		Action:   "create",
		UserID:   &getUserID,
		Data:     fmt.Sprint(category),
		IsActive: true,
	})
	return category.CategoryResponse(), err
}

func (s *Service) List(ctx context.Context, req *domain.ListCategoryRequest) ([]domain.CategoryResponse, int64, error) {
	cr := []domain.CategoryResponse{}
	categories, count, err := s.repo.List(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, category := range categories {
		cr = append(cr, *category.CategoryResponse())
	}
	return cr, count, nil
}

func (s *Service) Get(ctx context.Context, id string) (*domain.CategoryResponse, error) {
	category, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return category.CategoryResponse(), err
}

func (s *Service) Update(ctx context.Context, id string, req *domain.CategoryUpdateRequest) (*domain.CategoryResponse, error) {
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	mp := req.NewUpdateRequest()
	err = s.repo.Update(ctx, id, mp)
	if err != nil {
		return nil, err
	}
	category, err := s.repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Updated %s Category", category.Name),
		Action:   "Update",
		UserID:   &getUserID,
		Data:     fmt.Sprint(category),
		IsActive: true,
	})
	return category.CategoryResponse(), err
}

func (s *Service) Delete(ctx context.Context, id string) error {
	getUserID, err := getUserID(ctx)
	if err != nil {
		return err
	}
	category, err := s.repo.Get(ctx, id)
	if err != nil {
		return err
	}
	err = s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Created %s Category", category.Name),
		UserID:   &getUserID,
		Action:   "create",
		Data:     fmt.Sprint(category),
		IsActive: true,
	})
	return nil
}
