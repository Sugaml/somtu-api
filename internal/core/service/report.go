package service

import "github.com/sugaml/lms-api/internal/core/domain"

func (s *Service) GetLibraryDashboardStats() (*domain.LibraryDashboardStats, error) {
	result, err := s.repo.GetLibraryDashboardStats()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetMonthlyChartData() ([]domain.ChartData, error) {
	result, err := s.repo.GetMonthlyChartData()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (s *Service) GetDailyChartData(req *domain.ChartRequest) ([]domain.ChartData, error) {
	result, err := s.repo.GetDailyChartData(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetBorrowedBookStats() (*domain.BorrowedBookStats, error) {
	result, err := s.repo.GetBorrowedBookStats()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetBookProgramstats() (*[]domain.BookProgramstats, error) {
	result, err := s.repo.GetBookProgramstats()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *Service) GetInventorystats() (*domain.InventoryStats, error) {
	result, err := s.repo.GetInventorystats()
	if err != nil {
		return nil, err
	}
	return result, nil
}
