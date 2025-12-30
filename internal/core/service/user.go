package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
	util "github.com/sugaml/lms-api/internal/core/utils"
)

func (s *Service) CreateBulkUser(data *[]domain.UserRequest) ([]*domain.UserResponse, error) {
	var responses []*domain.UserResponse
	for _, req := range *data {
		err := req.Validate()
		if err != nil {
			return nil, err
		}
		data := domain.Convert[domain.UserRequest, domain.User](&req)
		data.Password, err = util.HashPassword(data.Password)
		if err != nil {
			return nil, err
		}
		if data.Role == "student" {
			data.Username = strings.ToLower(data.Username) + "." + data.StudentID
			studentExist, err := s.repo.GetStudentbyID(data.StudentID)
			if studentExist != nil && err == nil {
				return nil, errors.New("student already exist")
			}
		}
		result, err := s.repo.CreateUser(data)
		if err != nil {
			return nil, err
		}
		s.repo.CreateNotification(&domain.Notification{
			Title:    fmt.Sprintf("New student %s created.", result.Username),
			UserID:   result.ID,
			Module:   "user",
			Action:   "create",
			IsActive: true,
		})
		_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
			Title:    fmt.Sprintf("Created new student %s.", result.Username),
			UserID:   &result.ID,
			Action:   "create",
			Data:     string(domain.ConvertToJson(result)),
			IsActive: true,
		})
		logrus.Infof("Student %s created successfully", result.Username)
		responses = append(responses, domain.Convert[domain.User, domain.UserResponse](result))
	}
	return responses, nil
}

// CreateUser creates a new User
func (s *Service) CreateUser(req *domain.UserRequest) (*domain.UserResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.UserRequest, domain.User](req)
	data.Password, err = util.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	if data.Role == "student" {
		data.Username = strings.ToLower(data.Username) + "." + data.StudentID
		studentExist, err := s.repo.GetStudentbyID(data.StudentID)
		if studentExist != nil && err == nil {
			return nil, errors.New("student already exist")
		}
	}
	result, err := s.repo.CreateUser(data)
	if err != nil {
		return nil, err
	}
	s.repo.CreateNotification(&domain.Notification{
		Title:    fmt.Sprintf("New User %s created.", result.Username),
		UserID:   result.ID,
		Module:   "user",
		Action:   "create",
		IsActive: true,
	})
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Created new User %s.", result.Username),
		UserID:   &result.ID,
		Action:   "create",
		Data:     string(domain.ConvertToJson(result)),
		IsActive: true,
	})
	return domain.Convert[domain.User, domain.UserResponse](result), nil
}

func (s *Service) LoginUser(req *domain.LoginRequest) (*domain.LoginUserResponse, error) {
	user, err := s.repo.GetUserbyUsername(req.Username)
	if err != nil {
		return nil, err
	}
	err = util.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, err
	}
	accessToken, err := s.tokenMaker.CreateToken(
		user.Username,
		user.ID,
	)
	if err != nil {
		return nil, err
	}
	logrus.Infof("User %s logged in successfully", user.Role)
	s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("User %s logged in.", user.Username),
		UserID:   &user.ID,
		Action:   "login",
		Data:     fmt.Sprintf("User %s logged in successfully", user.Username),
		IsActive: true,
	})
	return &domain.LoginUserResponse{
		AccessToken: accessToken,
		User:        domain.Convert[domain.User, domain.UserResponse](user),
	}, nil
}

// ListUsers retrieves a list of Users
func (s *Service) ListUser(req *domain.UserListRequest) ([]*domain.UserResponse, int64, error) {
	var datas = []*domain.UserResponse{}
	results, count, err := s.repo.ListUser(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		data := domain.Convert[domain.User, domain.UserResponse](result)
		datas = append(datas, data)
	}
	return datas, count, nil
}

// ListUsers retrieves a list of Users
func (s *Service) ListStudent(req *domain.UserListRequest) ([]*domain.StudentResponse, int64, error) {
	var datas = []*domain.StudentResponse{}
	results, count, err := s.repo.ListStudent(req)
	if err != nil {
		return nil, count, err
	}
	for _, result := range results {
		data := domain.Convert[domain.User, domain.StudentResponse](result)
		if result.IsActive {
			data.Status = "warning"
		} else {
			data.Status = "overdue"
		}
		datas = append(datas, data)
	}
	return datas, count, nil
}

func (s *Service) GetUser(id string) (*domain.UserResponse, error) {
	result, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	data := domain.Convert[domain.User, domain.UserResponse](result)
	return data, nil
}

func (s *Service) UpdateUser(id string, req *domain.UserUpdateRequest) (*domain.UserResponse, error) {
	if id == "" {
		return nil, errors.New("required User id")
	}
	_, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	// update
	mp := req.NewUpdate()
	result, err := s.repo.UpdateUser(id, mp)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Updated %s User details.", result.Username),
		UserID:   &result.ID,
		Action:   "update",
		Data:     fmt.Sprint(req),
		IsActive: true,
	})
	data := domain.Convert[domain.User, domain.UserResponse](result)
	return data, nil
}

func (s *Service) DeleteUser(id string) (*domain.UserResponse, error) {
	result, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	if result.Role == "admin" {
		return nil, errors.New("cannot delete admin user")
	}
	if result.Role == "student" {
		CountBorrwedCopiesUserID, err := s.repo.CountBorrwedCopiesUserID(id)
		if err != nil {
			return nil, err
		}
		if CountBorrwedCopiesUserID > 0 {
			return nil, fmt.Errorf("user has %d copies borrowed cannot delete it", CountBorrwedCopiesUserID)
		}
	}
	err = s.repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	_, _ = s.repo.CreateAuditLog(&domain.AuditLog{
		Title:    fmt.Sprintf("Deleted %s parking area.", result.Username),
		UserID:   &result.ID,
		Action:   "delete",
		Data:     fmt.Sprint(result),
		IsActive: true,
	})
	return domain.Convert[domain.User, domain.UserResponse](result), nil
}
