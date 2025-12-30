package service

import (
	"context"
	"fmt"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (s *Service) CreateProgram(ctx context.Context, req *domain.ProgramRequest) (*domain.ProgramResponse, error) {
	data := &domain.Program{}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	data.NewProgram(req)
	result, err := s.repo.CreateProgram(ctx, data)
	if err != nil {
		return nil, err
	}
	userID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	s.repo.CreateNotification(&domain.Notification{
		Title:    "New program created: " + result.Name,
		UserID:   userID,
		Module:   "program",
		Action:   "create",
		IsActive: true,
	})
	response := domain.Convert[domain.Program, domain.ProgramResponse](result)
	return response, err
}

func (s *Service) LisProgram(ctx context.Context, req *domain.ListProgramRequest) ([]domain.ProgramResponse, int64, error) {
	cr := []domain.ProgramResponse{}
	categories, count, err := s.repo.ListProgram(ctx, req)
	if err != nil {
		return nil, count, err
	}
	for _, Program := range categories {
		cr = append(cr, *Program.ProgramResponse())
	}
	return cr, count, nil
}

func (s *Service) GetProgram(ctx context.Context, id string) (*domain.ProgramResponse, error) {
	Program, err := s.repo.GetProgram(ctx, id)
	if err != nil {
		return nil, err
	}
	return Program.ProgramResponse(), err
}

func (s *Service) UpdateProgram(ctx context.Context, id string, req *domain.ProgramUpdateRequest) (*domain.ProgramResponse, error) {
	mp := req.NewUpdateRequest()
	err := s.repo.UpdateProgram(ctx, id, mp)
	if err != nil {
		return nil, err
	}
	getUserID, err := getUserID(ctx)
	if err != nil {
		return nil, err
	}
	Program, err := s.repo.GetProgram(ctx, id)
	if err != nil {
		return nil, err
	}
	s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    "Program updated",
		Action:   "update",
		UserID:   &getUserID,
		Data:     fmt.Sprint(id),
		IsActive: true,
	})
	return Program.ProgramResponse(), err
}

func (s *Service) DeleteProgram(ctx context.Context, id string) error {
	getUserID, err := getUserID(ctx)
	if err != nil {
		return err
	}
	err = s.repo.DeleteProgram(ctx, id)
	if err != nil {
		return err
	}
	s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    "Program deleted",
		Action:   "delete",
		UserID:   &getUserID,
		Data:     fmt.Sprint(id),
		IsActive: true,
	})
	return nil
}
