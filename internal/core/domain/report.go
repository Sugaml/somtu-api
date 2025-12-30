package domain

type LibraryDashboardStats struct {
	ActiveStudents  int64 `json:"activeStudents"`
	AvailableBooks  int64 `json:"availableBooks"`
	BorrowedBooks   int64 `json:"borrowedBooks"`
	OverdueBooks    int64 `json:"overdueBooks"`
	PendingRequests int64 `json:"pendingRequests"`
	TotalBooks      int64 `json:"totalBooks"`
	TotalFines      int64 `json:"totalFines"`
	TotalStudents   int64 `json:"totalStudents"`
}

type ChartRequest struct {
	StartDate string `form:"start_date"`
	EndDate   string `form:"end_date"`
	Range     string `form:"range"`
}

type ChartData struct {
	Month         string `json:"month"`
	Date          string `json:"date"`
	Borrowed      int    `json:"borrowed"`
	Returned      int    `json:"returned"`
	Due           int    `json:"due"`
	Requests      int    `json:"requests"`
	TotalStudents int    `json:"totalStudents"`
	BooksAdded    int    `json:"booksAdded"`
}

type BorrowedBookStats struct {
	TotalBorrowedBooks int64 `json:"totalBorrowedBooks"`
	TotalOverdueBooks  int64 `json:"totalOverdueBooks"`
	PendingRequests    int64 `json:"pendingRequests"`
	DueSoon            int64 `json:"dueSoon"`
}

type BookProgramstats struct {
	ProgramID   string `json:"program_id"`
	ProgramName string `json:"program_name"`
	Count       int    `json:"count"`
}

type InventoryStats struct {
	TotalBooks      int64 `json:"totalBooks"`
	AvailableBooks  int64 `json:"availableBooks"`
	BorrowedBooks   int64 `json:"borrowedBooks"`
	OverdueBooks    int64 `json:"overdueBooks"`
	TotalStudents   int64 `json:"totalStudents"`
	ActiveStudents  int64 `json:"activeStudents"`
	PendingRequests int64 `json:"pendingRequests"`
	TotalFines      int64 `json:"totalFines"`
}

type DashboardStats struct {
	TotalBooks      int `json:"totalBooks"`
	AvailableBooks  int `json:"availableBooks"`
	BorrowedBooks   int `json:"borrowedBooks"`
	OverdueBooks    int `json:"overdueBooks"`
	TotalStudents   int `json:"totalStudents"`
	ActiveStudents  int `json:"activeStudents"`
	PendingRequests int `json:"pendingRequests"`
	TotalFines      int `json:"totalFines"`
}
