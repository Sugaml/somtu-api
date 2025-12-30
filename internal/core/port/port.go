package port

type Repository interface {
	AuditLogRepository
	UserRepository
	CategoryRepository
	ProgramRepository
	BookRepository
	BookCopyRepository
	FineRepository
	BorrowRepository
	ReportRepository
	NotificationRepository
}
type Service interface {
	AuditLogService
	UserService
	CategoryService
	ProgramService
	BookService
	BookCopyService
	FineService
	BorrowService
	ReportService
	NotificationService
}
