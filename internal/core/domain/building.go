package domain

type Building struct {
	BaseModel
	CampusID uint
	Campus   Campus

	Name string `gorm:"size:100;not null"`
	Code string `gorm:"size:10;not null;unique"`
}
