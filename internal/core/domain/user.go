// domain/user.go
package domain

import (
	"errors"
	"time"
)

type User struct {
	BaseModel
	Username       string `gorm:"unique;not null" json:"username"`
	Password       string `gorm:"not null" json:"password"`
	Dob            string `json:"dob"`
	MobileNumber   string `json:"mobile_number"`
	EnrollmentYear string `json:"enrollment_year"`
	Role           string `gorm:"not null" json:"role"` // student/librarian
	Email          string `gorm:"not null" json:"email"`
	Gender         string `json:"gender"`
	Level          string `json:"level"`
	Batch          string `json:"batch"`
	Section        string `json:"section"`
	Image          string `json:"image"`
	FullName       string `gorm:"column:full_name;not null" json:"full_name"`
	ProgramID      string `json:"program_id"`
	Program        string `json:"program"`
	Semester       string `json:"semester"`
	StudentID      string `gorm:"column:student_id" json:"student_id"`
	IsActive       bool   `gorm:"column:is_active;default:false" json:"is_active"`
}

type UserRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Dob            string `json:"dob"`
	MobileNumber   string `json:"mobile_number"`
	EnrollmentYear string `json:"enrollment_year"`
	Role           string `json:"role"`
	Email          string `json:"email"`
	Gender         string `json:"gender"`
	Level          string `json:"level"`
	Batch          string `json:"batch"`
	Section        string `json:"section"`
	Image          string `json:"image"`
	FullName       string `json:"full_name"`
	ProgramID      string `json:"program_id"`
	Program        string `json:"program"`
	Semester       string `json:"semester"`
	StudentID      string `json:"student_id"`
}

type UserListRequest struct {
	ListRequest
	Username       string `form:"username"`
	Password       string `form:"password"`
	Dob            string `form:"dob"`
	MobileNumber   string `form:"mobile_number"`
	Gender         string `json:"gender"`
	Level          string `json:"level"`
	Batch          string `json:"batch"`
	Section        string `json:"section"`
	EnrollmentYear string `form:"enrollment_year"`
	Role           string `form:"role"`
	Email          string `form:"email"`
	FullName       string `form:"full_name"`
	Image          string `json:"image"`
	Program        string `form:"program"`
	StudentID      string `form:"student_id"`
}

type UserAllUpdateRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Dob            string `json:"dob"`
	MobileNumber   string `json:"mobile_number"`
	Gender         string `json:"gender"`
	Level          string `json:"level"`
	Batch          string `json:"batch"`
	Section        string `json:"section"`
	EnrollmentYear string `json:"enrollment_year"`
	Role           string `json:"role"`
	Email          string `json:"email"`
	FullName       string `json:"full_name"`
	Image          string `json:"image"`
	Program        string `json:"program"`
	StudentID      string `json:"student_id"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type LoginUserResponse struct {
	AccessToken string `json:"access_token"`
	User        *UserResponse
}

type UserUpdateRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	Dob            string `json:"dob"`
	Gender         string `json:"gender"`
	Level          string `json:"level"`
	Batch          string `json:"batch"`
	Section        string `json:"section"`
	MobileNumber   string `json:"mobile_number"`
	EnrollmentYear string `json:"enrollment_year"`
	Role           string `json:"role"`
	Image          string `json:"image"`
	Email          string `json:"email"`
	FullName       string `json:"full_name"`
	Program        string `json:"program"`
	StudentID      string `json:"student_id"`
}

type UserResponse struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	Username       string    `json:"username"`
	Dob            string    `json:"dob"`
	Gender         string    `json:"gender"`
	Level          string    `json:"level"`
	Batch          string    `json:"batch"`
	Section        string    `json:"section"`
	MobileNumber   string    `json:"mobile_number"`
	EnrollmentYear string    `json:"enrollment_year"`
	Role           string    `json:"role"`
	Email          string    `json:"email"`
	Image          string    `json:"image"`
	FullName       string    `json:"full_name"`
	Program        string    `json:"program"`
	Semester       string    `json:"semester"`
	StudentID      string    `json:"student_id"`
	IsActive       bool      `json:"is_active"`
}

type StudentResponse struct {
	ID             string `json:"id"`
	FullName       string `json:"full_name"`
	Dob            string `json:"dob"`
	Role           string `json:"role"`
	Gender         string `json:"gender"`
	Level          string `json:"level"`
	Batch          string `json:"batch"`
	Section        string `json:"section"`
	MobileNumber   string `json:"mobile_number"`
	EnrollmentYear string `json:"enrollment_year"`
	Email          string `json:"email"`
	StudentID      string `json:"student_id"`
	Program        string `json:"program"`
	Semester       string `json:"semester"`
	Image          string `json:"image"`
	BorrowedCount  int    `json:"borrowed_count" default:"10"`
	OverdueCount   int    `json:"overdue_count"  default:"10"`
	Fines          int    `json:"fines"  default:"10"`
	Status         string `json:"status"  default:"clearance"`
	ProfileImage   string `json:"profile_image"`
}

func (u *UserRequest) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if u.Role == "" {
		return errors.New("role is required")
	}
	// if u.Email == "" {
	// 	return errors.New("email is required")
	// }
	if u.FullName == "" {
		return errors.New("full name is required")
	}
	if u.Role == "Student" {
		if u.Program == "" {
			return errors.New("program is required")
		}
		if u.StudentID == "" {
			return errors.New("student id is required")
		}
	}
	return nil
}

func (r *UserUpdateRequest) NewUpdate() Map {
	mp := map[string]interface{}{}
	if r.Username != "" {
		mp["username"] = r.Username
	}
	if r.Image != "" {
		mp["image"] = r.Image
	}
	if r.Password != "" {
		mp["password"] = r.Password
	}
	if r.Role != "" {
		mp["role"] = r.Role
	}
	if r.Dob != "" {
		mp["dob"] = r.Dob
	}
	if r.Gender != "" {
		mp["gender"] = r.Gender
	}
	if r.Level != "" {
		mp["level"] = r.Level
	}
	if r.Batch != "" {
		mp["batch"] = r.Batch
	}
	if r.Section != "" {
		mp["section"] = r.Section
	}
	if r.MobileNumber != "" {
		mp["mobile_number"] = r.MobileNumber
	}
	if r.EnrollmentYear != "" {
		mp["enrollment_year"] = r.EnrollmentYear
	}
	if r.Email != "" {
		mp["email"] = r.Email
	}
	if r.FullName != "" {
		mp["full_name"] = r.FullName
	}
	if r.Program != "" {
		mp["program"] = r.Program
	}
	return mp
}
