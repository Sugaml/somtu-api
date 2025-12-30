package domain

type Role struct {
	BaseModel
	Name string `gorm:"unique;not null"` // STUDENT, TEACHER, ADMIN, DIRECTOR
}

type UserRole struct {
	UserID uint
	RoleID uint
}
