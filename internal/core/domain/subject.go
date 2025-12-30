package domain

type Subject struct {
	BaseModel
	ProgramID uint
	Program   Program

	Code        string `gorm:"size:20;unique"`
	Name        string `gorm:"size:100;not null"`
	CreditHours int
	IsLab       bool
}
