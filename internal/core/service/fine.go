package service

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

// CreateFine creates a new Fine
func (s *Service) CreateFine(req *domain.FineRequest) (*domain.FineResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.FineRequest, domain.Fine](req)
	result, err := s.repo.CreateFine(data)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.Fine, domain.FineResponse](result), nil
}

// ListFines retrieves a list of Fines
func (s *Service) ListFine(req *domain.ListFineRequest) ([]*domain.FineResponse, int64, error) {
	var datas = []*domain.FineResponse{}
	results, count, err := s.repo.ListFine(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		data := domain.Convert[domain.Fine, domain.FineResponse](result)
		datas = append(datas, data)
	}
	return datas, count, nil
}

func (s *Service) GetFine(id string) (*domain.FineResponse, error) {
	result, err := s.repo.GetFine(id)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.Fine, domain.FineResponse](result)
	return data, nil
}

func (s *Service) UpdateFine(id string, req *domain.UpdateFineRequest) (*domain.FineResponse, error) {
	if id == "" {
		return nil, errors.New("required Fine id")
	}
	_, err := s.repo.GetFine(id)
	if err != nil {
		return nil, err
	}

	// update
	mp := req.NewUpdate()
	result, err := s.repo.UpdateFine(id, mp)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.Fine, domain.FineResponse](result)
	return data, nil
}

func (s *Service) DeleteFine(id string) (*domain.FineResponse, error) {
	result, err := s.repo.GetFine(id)
	if err != nil {
		return nil, err
	}
	err = s.repo.DeleteFine(id)
	if err != nil {
		return nil, err
	}
	return domain.Convert[domain.Fine, domain.FineResponse](result), nil
}
