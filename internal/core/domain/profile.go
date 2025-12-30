package domain

type StudentProfile struct {
	BaseModel

	UserID uint
	User   User

	StudentID      string `gorm:"unique"`
	EnrollmentYear string
	Batch          string
	Section        string

	ProgramID uint
	Program   Program

	SemesterID uint
	Semester   Semester
}

type TeacherProfile struct {
	BaseModel

	UserID uint
	User   User

	EmployeeID  string `gorm:"unique"`
	Designation string // Lecturer, Assistant Prof
	FacultyID   uint
	Faculty     Faculty
}

type StaffProfile struct {
	BaseModel

	UserID uint
	User   User

	EmployeeID string `gorm:"unique"`
	Position   string // ADMIN, DIRECTOR, DEPUTY_DIRECTOR
	Office     string
}
