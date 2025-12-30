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
	FloorID Floor
	Floor   Floor

	RoomNumber   string   `gorm:"size:20"`
	RoomCode     string   `gorm:"size:20;unique"`
	RoomType     RoomType `gorm:"type:enum('CLASSROOM','LAB','SEMINAR','AUDITORIUM')"`
	Capacity     int
	HasProjector bool
	HasAC        bool
	Status       string `gorm:"default:ACTIVE"`
}
