package service

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateAuditLog creates a new AuditLog
func (s *Service) CreateAuditLog(req *domain.AuditLogRequest) (*domain.AuditLogResponse, error) {
	data := domain.Convert[domain.AuditLogRequest, domain.AuditLog](req)
	err := data.Validate()
	if err != nil {
		return nil, err
	}
	result, err := s.repo.CreateAuditLog(data)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.AuditLog, domain.AuditLogResponse](result), nil
}

// ListAuditLogs retrieves a list of AuditLogs
func (s *Service) ListAuditLog(req *domain.ListAuditLogRequest) ([]*domain.AuditLogResponse, int64, error) {
	var datas = []*domain.AuditLogResponse{}
	req.Size = 5
	results, count, err := s.repo.ListAuditLog(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		datas = append(datas, domain.Convert[domain.AuditLog, domain.AuditLogResponse](result))
	}
	return datas, count, nil
}

func (s *Service) GetAuditLog(id string) (*domain.AuditLogResponse, error) {
	result, err := s.repo.GetAuditLog(id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.AuditLog, domain.AuditLogResponse](result), nil
}

func (s *Service) UpdateAuditLog(id string, req *domain.AuditLogUpdateRequest) (*domain.AuditLogResponse, error) {
	if id == "" {
		return nil, errors.New("required AuditLog id")
	}
	_, err := s.repo.GetAuditLog(id)
	if err != nil {
		return nil, err
	}
	mp := req.NewUpdate()
	result, err := s.repo.UpdateAuditLog(id, mp)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.AuditLog, domain.AuditLogResponse](result), nil
}

func (s *Service) DeleteAuditLog(id string) error {
	if id == "" {
		return errors.New("required AuditLog id")
	}
	err := s.repo.DeleteAuditLog(id)
	if err != nil {
		return err
	}
	return nil
}
