package domain

type Notification struct {
	BaseModel
	UserID      string `gorm:"not null" json:"user_id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Module      string `gorm:"not null" json:"module"` // 'book' | 'program' | 'user' | 'general'
	Action      string `gorm:"not null" json:"action"` // 'create' | 'update' | 'delete'
	Type        string `gorm:"not null" json:"type"`
	IsRead      bool   `gorm:"column:is_read;default:false" json:"is_read"`
	IsActive    bool   `gorm:"column:is_active;default:false" json:"is_active"`
}

type NotificationRequest struct {
	BaseModel
	UserID      string `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Type        string `json:"type" validate:"required"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool   `json:"is_read"`
}

type UpdateNotificationRequest struct {
	BaseModel
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool   `json:"is_read"`
}

type ListNotificationRequest struct {
	ListRequest
	UserID      string `form:"user_id"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Type        string `form:"type"` // 'due_reminder' | 'fine' | 'request_approved' | 'general'
	IsRead      bool   `form:"is_read"`
}

type NotificationResponse struct {
	BaseModel
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Module      string `json:"module"`
	Action      string `json:"action"`
	Type        string `json:"type"`
	IsRead      bool   `json:"is_read"`
}

func (u *NotificationRequest) Validate() error {
	return nil
}

func (r *UpdateNotificationRequest) NewUpdate() Map {
	return nil
}
