package port

import "github.com/sugaml/lms-api/internal/core/domain"

// type AuditLogRepository interface is an interface for interacting with type AuditLog-related data
type AuditLogRepository interface {
	CreateAuditLog(data *domain.AuditLog) (*domain.AuditLog, error)
	ListAuditLog(req *domain.ListAuditLogRequest) ([]*domain.AuditLog, int64, error)
	GetAuditLog(id string) (*domain.AuditLog, error)
	UpdateAuditLog(id string, req domain.Map) (*domain.AuditLog, error)
	DeleteAuditLog(id string) error
}

// type AuditLogService interface is an interface for interacting with type AuditLog-related data
type AuditLogService interface {
	CreateAuditLog(data *domain.AuditLogRequest) (*domain.AuditLogResponse, error)
	ListAuditLog(req *domain.ListAuditLogRequest) ([]*domain.AuditLogResponse, int64, error)
	GetAuditLog(id string) (*domain.AuditLogResponse, error)
	UpdateAuditLog(id string, req *domain.AuditLogUpdateRequest) (*domain.AuditLogResponse, error)
	DeleteAuditLog(id string) error
}
