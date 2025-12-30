package domain

import "time"

type TimeSlot struct {
	BaseModel
	StartTime time.Time
	EndTime   time.Time
}
