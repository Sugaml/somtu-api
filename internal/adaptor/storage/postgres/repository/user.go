package repository

import (
	"errors"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) CreateUser(data *domain.User) (*domain.User, error) {
	if err := r.db.Model(&domain.User{}).Create(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) ListUser(req *domain.UserListRequest) ([]*domain.User, int64, error) {
	var datas []*domain.User
	var count int64
	f := r.db.Model(&domain.User{})
	if req.Query != "" {
		req.SortColumn = "score desc, " + req.SortColumn
	}
	err := f.Count(&count).
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, nil
}

func (r *Repository) ListStudent(req *domain.UserListRequest) ([]*domain.User, int64, error) {
	var datas []*domain.User
	var count int64
	f := r.db.Model(&domain.User{}).Where("role = ?", "student")
	if req.Query != "" {
		req.SortColumn = "score desc, " + req.SortColumn
	}
	if req.FullName != "" {
		f = f.Where("full_name ILIKE ?", "%"+req.FullName+"%")
	}
	if req.Program != "all" {
		f = f.Where("program ILIKE ?", "%"+req.Program+"%")
	}
	if req.Dob != "" {
		f = f.Where("dob = ?", req.Dob)
	}
	if req.StudentID != "" {
		f = f.Where("student_id = ?", req.StudentID)
	}
	if req.Gender != "" {
		f = f.Where("gender = ?", req.Gender)
	}
	if req.Username != "" {
		f = f.Where("username = ?", req.Username)
	}
	err := f.Count(&count).
		Order(req.SortColumn + " " + req.SortDirection).
		Limit(req.Size).
		Offset(req.Size * (req.Page - 1)).
		Find(&datas).Error
	if err != nil {
		return nil, count, err
	}
	return datas, count, nil
}

func (r *Repository) GetUser(id string) (*domain.User, error) {
	var data domain.User
	if err := r.db.Model(&domain.User{}).
		Take(&data, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) GetStudentbyID(studentID string) (*domain.User, error) {
	var data domain.User
	if err := r.db.Model(&domain.User{}).
		Take(&data, "student_id = ? and role = ?", studentID, "student").Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) GetUserbyUsername(username string) (*domain.User, error) {
	var data domain.User
	if err := r.db.Model(&domain.User{}).
		Take(&data, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *Repository) UpdateUser(id string, req domain.Map) (*domain.User, error) {
	if id == "" {
		return nil, errors.New("required user id")
	}
	data := &domain.User{}
	err := r.db.Model(&domain.User{}).Where("id = ?", id).Updates(req.ToMap()).Take(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) DeleteUser(id string) error {
	return r.db.Model(&domain.User{}).Where("id = ?", id).Delete(&domain.User{}).Error
}
