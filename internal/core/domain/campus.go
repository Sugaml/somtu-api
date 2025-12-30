package domain

type Campus struct {
	BaseModel
	Name     string `gorm:"size:100;not null;unique"`
	Location string `gorm:"size:255"`
}
