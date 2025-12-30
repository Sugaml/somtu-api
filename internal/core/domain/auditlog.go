package domain

import (
	"time"
)

// Define the AuditLog struct with JSON tags
type AuditLog struct {
	BaseModel
	Title       string  `json:"title"`
	UserID      *string `json:"user_id"`
	Action      string  `json:"action"`       // The action performed (e.g., "CREATE", "UPDATE", "DELETE")
	PerformedBy string  `json:"performed_by"` // The user or system that performed the action
	Data        string  `json:"data"`
	Details     string  `json:"details"` // Additional details about the action (optional)
	Remarks     string  `json:"remarks"`
	IsActive    bool    `json:"is_active"`
}

type ListAuditLogRequest struct {
	ListRequest
	UserID      string `form:"user_id"`
	Action      string `form:"action"`
	PerformedBy string `form:"performed_by"`
}

type AuditLogUpdateRequest struct {
	Action      string `json:"action"`
	PerformedBy string `json:"performed_by"`
	Remarks     string `json:"remarks"`
}

// Define the AuditLogRequest struct with JSON tags
type AuditLogRequest struct {
	Title       string `json:"title"`
	UserID      string `form:"user_id" binding:"omitempty" swaggerignore:"true"`
	Action      string `json:"action"`       // The action performed (e.g., "CREATE", "UPDATE", "DELETE")
	PerformedBy string `json:"performed_by"` // The user or system that performed the action
	Data        string `json:"data"`
	Details     string `json:"details"` // Additional details about the action (optional)
	Remarks     string `json:"remarks"`
}

// Define the AuditLogResponse struct with JSON tags
type AuditLogResponse struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Module      string    `json:"module,omitempty"`
	Title       string    `json:"title,omitempty"`
	UserID      string    `json:"user_id,omitempty"`
	Action      string    `json:"action,omitempty"`       // The action performed (e.g., "CREATE", "UPDATE", "DELETE")
	PerformedBy string    `json:"performed_by,omitempty"` // The user or system that performed the action
	Data        string    `json:"data,omitempty"`
	Details     string    `json:"details,omitempty"` // Additional details about the action (optional)
	Remarks     string    `json:"remarks,omitempty"`
	IsActive    bool      `json:"is_active,omitempty"`
}

func (a *AuditLogUpdateRequest) NewUpdate() Map {
	return nil
}

func (a *AuditLog) Validate() error {
	return nil
}
