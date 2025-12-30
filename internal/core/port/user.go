package port

import (
	"github.com/sugaml/lms-api/internal/core/domain"
)

// type UserRepository interface is an interface for interacting with type Announcement-related data
type UserRepository interface {
	CreateUser(data *domain.User) (*domain.User, error)
	ListUser(req *domain.UserListRequest) ([]*domain.User, int64, error)
	ListStudent(req *domain.UserListRequest) ([]*domain.User, int64, error)
	GetUser(id string) (*domain.User, error)
	GetStudentbyID(studentID string) (*domain.User, error)
	GetUserbyUsername(username string) (*domain.User, error)
	UpdateUser(id string, req domain.Map) (*domain.User, error)
	DeleteUser(id string) error
}

// type UserService interface is an interface for interacting with type Announcement-related data
type UserService interface {
	CreateUser(data *domain.UserRequest) (*domain.UserResponse, error)
	CreateBulkUser(data *[]domain.UserRequest) ([]*domain.UserResponse, error)
	LoginUser(req *domain.LoginRequest) (*domain.LoginUserResponse, error)
	ListUser(req *domain.UserListRequest) ([]*domain.UserResponse, int64, error)
	ListStudent(req *domain.UserListRequest) ([]*domain.StudentResponse, int64, error)
	GetUser(id string) (*domain.UserResponse, error)
	UpdateUser(id string, req *domain.UserUpdateRequest) (*domain.UserResponse, error)
	DeleteUser(id string) (*domain.UserResponse, error)
}
