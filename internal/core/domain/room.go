package domain

type RoomType string

const (
	ClassRoom   RoomType = "CLASSROOM"
	LabRoom     RoomType = "LAB"
	SeminarRoom RoomType = "SEMINAR"
	Auditorium  RoomType = "AUDITORIUM"
)

type Room struct {
	BaseModel

	FloorID uint
	Floor   Floor

	RoomNumber   string   `gorm:"size:20"`
	RoomCode     string   `gorm:"size:20;unique"`
	RoomType     RoomType `gorm:"type:varchar(20);not null"`
	Capacity     int
	HasProjector bool
	HasAC        bool
	Status       string `gorm:"default:'ACTIVE'"`
}
