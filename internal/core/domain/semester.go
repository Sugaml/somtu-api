package domain

type Semester struct {
	BaseModel
	ProgramID uint
	Program   Program

	Name       string `gorm:"size:50"`
	SemesterNo int
}
