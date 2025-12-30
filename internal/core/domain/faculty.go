package domain

type Faculty struct {
	BaseModel
	Name string `gorm:"size:100;not null;unique"`
}
