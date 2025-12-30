package domain

type Floor struct {
	BaseModel
	BuildingID uint
	Building   Building

	FloorNumber int
	Description string `gorm:"size:100"`
}
