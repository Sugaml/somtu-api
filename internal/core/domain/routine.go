package domain

type DayOfWeek string

const (
	Monday    DayOfWeek = "MON"
	Tuesday   DayOfWeek = "TUE"
	Wednesday DayOfWeek = "WED"
	Thursday  DayOfWeek = "THU"
	Friday    DayOfWeek = "FRI"
	Saturday  DayOfWeek = "SAT"
)

type ClassRoutine struct {
	BaseModel

	FacultyID uint
	Faculty   Faculty

	ProgramID uint
	Program   Program

	SemesterID uint
	Semester   Semester

	SubjectID uint
	Subject   Subject

	TeacherID uint // from user table

	RoomID uint
	Room   Room

	TimeSlotID uint
	TimeSlot   TimeSlot

	DayOfWeek    DayOfWeek `gorm:"type:enum('MON','TUE','WED','THU','FRI','SAT')"`
	AcademicYear string    `gorm:"size:20"`
}
