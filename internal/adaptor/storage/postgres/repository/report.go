package repository

import (
	"fmt"
	"time"

	"github.com/sugaml/lms-api/internal/core/domain"
)

func (r *Repository) GetLibraryDashboardStats() (*domain.LibraryDashboardStats, error) {
	var stats domain.LibraryDashboardStats
	now := time.Now()

	//Count total students
	if err := r.db.Model(&domain.User{}).
		Where("role = ?", "student").
		Count(&stats.TotalStudents).Error; err != nil {
		return nil, err
	}

	//Count active students
	if err := r.db.Model(&domain.User{}).
		Where("role = ? AND is_active = ?", "student", true).
		Count(&stats.ActiveStudents).Error; err != nil {
		return nil, err
	}

	//Count total active books
	if err := r.db.Model(&domain.Book{}).
		Count(&stats.TotalBooks).Error; err != nil {
		return nil, err
	}

	// Count total pending books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND returned_date IS NULL AND is_active = ?", "borrowed", true).
		Count(&stats.PendingRequests).Error; err != nil {
		return nil, err
	}

	// Count total borrowed books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND due_date < ? AND returned_date IS NULL AND is_active = ?", "borrowed", now, true).
		Count(&stats.BorrowedBooks).Error; err != nil {
		return nil, err
	}

	// Count overdue books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND due_date < ? AND returned_date IS NULL AND is_active = ?", "borrowed", now, true).
		Count(&stats.OverdueBooks).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *Repository) GetMonthlyChartData() ([]domain.ChartData, error) {
	var results []domain.ChartData

	type TempData struct {
		Month         string
		YearMonth     string
		Borrowed      int
		Returned      int
		Due           int
		Requests      int
		TotalStudents int
		BooksAdded    int
	}

	var data []TempData

	err := r.db.Raw(`
			WITH borrow_summary AS (
				SELECT 
					TO_CHAR(borrowed_date, 'Mon') AS month,
					TO_CHAR(borrowed_date, 'YYYY-MM') AS year_month,
					COUNT(*) FILTER (WHERE status = 'borrowed') AS borrowed,
					COUNT(*) FILTER (WHERE status = 'returned') AS returned,
					COUNT(*) FILTER (WHERE status = 'overdue') AS due,
					COUNT(*) FILTER (WHERE status = 'pending') AS requests
				FROM borrowed_books
				GROUP BY year_month, month
			),
			student_summary AS (
				SELECT 
					TO_CHAR(created_at, 'YYYY-MM') AS year_month,
					COUNT(*) AS total_students
				FROM users
				WHERE role = 'student'
				GROUP BY year_month
			),
			book_summary AS (
				SELECT 
					TO_CHAR(created_at, 'YYYY-MM') AS year_month,
					COUNT(*) AS books_added
				FROM books
				GROUP BY year_month
			)
			SELECT 
				bs.month,
				bs.year_month,
				bs.borrowed,
				bs.returned,
				bs.due,
				bs.requests,
				COALESCE(ss.total_students, 0) AS total_students,
				COALESCE(bo.books_added, 0) AS books_added
			FROM borrow_summary bs
			LEFT JOIN student_summary ss ON ss.year_month = bs.year_month
			LEFT JOIN book_summary bo ON bo.year_month = bs.year_month
			ORDER BY bs.year_month ASC
		`).Scan(&data).Error

	if err != nil {
		return nil, err
	}

	for _, d := range data {
		results = append(results, domain.ChartData{
			Month:         d.Month,
			Date:          d.YearMonth,
			Borrowed:      d.Borrowed,
			Returned:      d.Returned,
			Due:           d.Due,
			Requests:      d.Requests,
			TotalStudents: d.TotalStudents,
			BooksAdded:    d.BooksAdded,
		})
	}

	return results, nil
}
func (r *Repository) GetDailyChartData(req *domain.ChartRequest) ([]domain.ChartData, error) {
	var startDate, endDate time.Time
	var err error

	// Parse provided dates if available
	if req.StartDate != "" {
		startDate, err = time.Parse("2006-01-02", req.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start date: %w", err)
		}
	}
	if req.EndDate != "" {
		endDate, err = time.Parse("2006-01-02", req.EndDate)
		if err != nil {
			return nil, fmt.Errorf("invalid end date: %w", err)
		}
		// Set end date to end of day for inclusive range
		endDate = endDate.Add(24*time.Hour - time.Second)
	}

	// Set defaults if not provided
	switch req.Range {
	case "daily":
		if req.StartDate == "" || req.EndDate == "" {
			endDate = time.Now()
			startDate = endDate.AddDate(0, 0, -7) // last 7 days
		}
	case "weekly":
		if req.StartDate == "" || req.EndDate == "" {
			endDate = time.Now()
			startDate = endDate.AddDate(0, 0, -30) // last ~4 weeks
		}
	case "monthly":
		if req.StartDate == "" || req.EndDate == "" {
			endDate = time.Now()
			startDate = endDate.AddDate(0, -12, 0) // last 12 months
		}
	case "quarterly":
		if req.StartDate == "" || req.EndDate == "" {
			endDate = time.Now()
			startDate = endDate.AddDate(0, -12, 0) // last 12 months (4 quarters)
		}
	case "yearly":
		if req.StartDate == "" || req.EndDate == "" {
			endDate = time.Now()
			startDate = endDate.AddDate(-5, 0, 0) // last 5 years
		}
	default:
		return nil, fmt.Errorf("unsupported range: %s", req.Range)
	}

	// Date formatting for SQL
	startDateStr := startDate.Format("2006-01-02")
	endDateStr := endDate.Format("2006-01-02")

	// Pick grouping based on range for each table
	borrowGroupExpr := "TO_CHAR(borrowed_date, 'YYYY-MM-DD')" // daily
	studentGroupExpr := "TO_CHAR(created_at, 'YYYY-MM-DD')"   // daily
	bookGroupExpr := "TO_CHAR(created_at, 'YYYY-MM-DD')"      // daily

	switch req.Range {
	case "weekly":
		borrowGroupExpr = "TO_CHAR(borrowed_date, 'IYYY-IW')" // ISO week
		studentGroupExpr = "TO_CHAR(created_at, 'IYYY-IW')"   // ISO week
		bookGroupExpr = "TO_CHAR(created_at, 'IYYY-IW')"      // ISO week
	case "monthly":
		borrowGroupExpr = "TO_CHAR(borrowed_date, 'YYYY-MM')"
		studentGroupExpr = "TO_CHAR(created_at, 'YYYY-MM')"
		bookGroupExpr = "TO_CHAR(created_at, 'YYYY-MM')"
	case "quarterly":
		borrowGroupExpr = "CONCAT(EXTRACT(YEAR FROM borrowed_date), '-Q', EXTRACT(QUARTER FROM borrowed_date))"
		studentGroupExpr = "CONCAT(EXTRACT(YEAR FROM created_at), '-Q', EXTRACT(QUARTER FROM created_at))"
		bookGroupExpr = "CONCAT(EXTRACT(YEAR FROM created_at), '-Q', EXTRACT(QUARTER FROM created_at))"
	case "yearly":
		borrowGroupExpr = "TO_CHAR(borrowed_date, 'YYYY')"
		studentGroupExpr = "TO_CHAR(created_at, 'YYYY')"
		bookGroupExpr = "TO_CHAR(created_at, 'YYYY')"
	}

	// SQL query: join three summaries (borrow, student, books)
	query := fmt.Sprintf(`
		WITH borrow_summary AS (
			SELECT %s AS grp,
			       COUNT(CASE WHEN status = 'borrowed' THEN 1 END) AS borrowed,
			       COUNT(CASE WHEN status = 'returned' THEN 1 END) AS returned,
			       COUNT(CASE WHEN status = 'overdue' THEN 1 END) AS due,
			       COUNT(CASE WHEN status = 'pending' THEN 1 END) AS requests
			FROM borrowed_books
			WHERE borrowed_date BETWEEN ? AND ?
			GROUP BY grp
		),
		student_summary AS (
			SELECT %s AS grp,
			       COUNT(*) AS total_students
			FROM users
			WHERE created_at BETWEEN ? AND ?
			GROUP BY grp
		),
		book_summary AS (
			SELECT %s AS grp,
			       COUNT(*) AS books_added
			FROM books
			WHERE created_at BETWEEN ? AND ?
			GROUP BY grp
		)
		SELECT COALESCE(b.grp, s.grp, bk.grp) AS grp,
		       COALESCE(b.borrowed, 0) AS borrowed,
		       COALESCE(b.returned, 0) AS returned,
		       COALESCE(b.due, 0) AS due,
		       COALESCE(b.requests, 0) AS requests,
		       COALESCE(s.total_students, 0) AS total_students,
		       COALESCE(bk.books_added, 0) AS books_added
		FROM borrow_summary b
		FULL OUTER JOIN student_summary s ON b.grp = s.grp
		FULL OUTER JOIN book_summary bk ON COALESCE(b.grp, s.grp) = bk.grp
	`, borrowGroupExpr, studentGroupExpr, bookGroupExpr)

	type Temp struct {
		Grp           string
		Borrowed      int
		Returned      int
		Due           int
		Requests      int
		TotalStudents int
		BooksAdded    int
	}
	var rawData []Temp

	err = r.db.Raw(
		query,
		startDateStr, endDateStr, // borrow
		startDateStr, endDateStr, // students
		startDateStr, endDateStr, // books
	).Scan(&rawData).Error
	if err != nil {
		return nil, err
	}

	// Build map for quick lookup
	dataMap := make(map[string]Temp)
	for _, row := range rawData {
		dataMap[row.Grp] = row
	}

	// Fill response
	var result []domain.ChartData

	switch req.Range {
	case "daily":
		for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 1) {
			key := formatGroupKey(d, req.Range)
			row := dataMap[key]

			result = append(result, domain.ChartData{
				Month:         d.Format("Mon 02"),
				Date:          d.Format("2006-01-02"),
				Borrowed:      row.Borrowed,
				Returned:      row.Returned,
				Due:           row.Due,
				Requests:      row.Requests,
				TotalStudents: row.TotalStudents,
				BooksAdded:    row.BooksAdded,
			})
		}
	case "weekly":
		for d := startDate; !d.After(endDate); d = d.AddDate(0, 0, 7) {
			key := formatGroupKey(d, req.Range)
			row := dataMap[key]

			year, week := d.ISOWeek()
			result = append(result, domain.ChartData{
				Month:         fmt.Sprintf("Wk %02d", week),
				Date:          fmt.Sprintf("%04d-W%02d", year, week),
				Borrowed:      row.Borrowed,
				Returned:      row.Returned,
				Due:           row.Due,
				Requests:      row.Requests,
				TotalStudents: row.TotalStudents,
				BooksAdded:    row.BooksAdded,
			})
		}
	case "monthly":
		for d := startDate; !d.After(endDate); d = d.AddDate(0, 1, 0) {
			key := formatGroupKey(d, req.Range)
			row := dataMap[key]

			result = append(result, domain.ChartData{
				Month:         d.Format("Jan 2006"),
				Date:          d.Format("2006-01"),
				Borrowed:      row.Borrowed,
				Returned:      row.Returned,
				Due:           row.Due,
				Requests:      row.Requests,
				TotalStudents: row.TotalStudents,
				BooksAdded:    row.BooksAdded,
			})
		}
	case "quarterly":
		for d := startDate; !d.After(endDate); d = d.AddDate(0, 3, 0) {
			key := formatGroupKey(d, req.Range)
			row := dataMap[key]

			quarter := ((d.Month() - 1) / 3) + 1
			result = append(result, domain.ChartData{
				Month:         fmt.Sprintf("Q%d %d", quarter, d.Year()),
				Date:          fmt.Sprintf("%d-Q%d", d.Year(), quarter),
				Borrowed:      row.Borrowed,
				Returned:      row.Returned,
				Due:           row.Due,
				Requests:      row.Requests,
				TotalStudents: row.TotalStudents,
				BooksAdded:    row.BooksAdded,
			})
		}
	case "yearly":
		for d := startDate; !d.After(endDate); d = d.AddDate(1, 0, 0) {
			key := formatGroupKey(d, req.Range)
			row := dataMap[key]

			result = append(result, domain.ChartData{
				Month:         d.Format("2006"),
				Date:          d.Format("2006"),
				Borrowed:      row.Borrowed,
				Returned:      row.Returned,
				Due:           row.Due,
				Requests:      row.Requests,
				TotalStudents: row.TotalStudents,
				BooksAdded:    row.BooksAdded,
			})
		}
	}

	return result, nil
}

// Helper to make Go keys match Postgres TO_CHAR
func formatGroupKey(d time.Time, rng string) string {
	switch rng {
	case "daily":
		return d.Format("2006-01-02")
	case "weekly":
		year, week := d.ISOWeek()
		return fmt.Sprintf("%04d-%02d", year, week)
	case "monthly":
		return d.Format("2006-01")
	case "quarterly":
		quarter := ((d.Month() - 1) / 3) + 1
		return fmt.Sprintf("%d-Q%d", d.Year(), quarter)
	case "yearly":
		return d.Format("2006")
	default:
		return d.Format("2006-01-02")
	}
}

func (r *Repository) GetBorrowedBookStats() (*domain.BorrowedBookStats, error) {
	var stats domain.BorrowedBookStats
	now := time.Now()

	// Count total borrowed books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Count(&stats.TotalBorrowedBooks).Error; err != nil {
		return nil, err
	}

	// Count overdue books
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND due_date < ? AND returned_date IS NULL AND is_active = ?", "borrowed", now, true).
		Count(&stats.TotalOverdueBooks).Error; err != nil {
		return nil, err
	}

	// Count pending requests
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND is_active = ?", "pending", true).
		Count(&stats.PendingRequests).Error; err != nil {
		return nil, err
	}

	// Count due soon (within 3 days)
	threeDaysLater := now.Add(72 * time.Hour)
	if err := r.db.Model(&domain.BorrowedBook{}).
		Where("status = ? AND due_date BETWEEN ? AND ? AND is_active = ?", "borrowed", now, threeDaysLater, true).
		Count(&stats.DueSoon).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *Repository) GetBookProgramstats() (*[]domain.BookProgramstats, error) {
	var stats []domain.BookProgramstats

	if err := r.db.Model(&domain.User{}).
		Select("program as program_name, count(*) as count").
		Where("role = ?", "student").
		Group("program").
		Scan(&stats).Error; err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *Repository) GetInventorystats() (*domain.InventoryStats, error) {
	var stats domain.InventoryStats
	var totalBooks int64
	var borrowedBooks int64
	var overdueBooks int64
	var totalStudents int64
	var activeStudents int64
	var pendingRequests int64
	var totalFines int64

	// Queries
	r.db.Model(&domain.Book{}).Count(&totalBooks)
	r.db.Model(&domain.BorrowedBook{}).Where("status = ?", "borrowed").Count(&borrowedBooks)
	r.db.Model(&domain.BorrowedBook{}).Where("status = ?", "overdue").Count(&overdueBooks)
	r.db.Model(&domain.User{}).Where("role = ?", "student").Count(&totalStudents)
	r.db.Model(&domain.User{}).Where("role = ? AND is_active = ?", "student", true).Count(&activeStudents)
	r.db.Model(&domain.BorrowedBook{}).Where("status = ?", "pending").Count(&pendingRequests)
	// r.db.Model(&domain.Fine{}).Where("status = ?", "pending").Select("SUM(amount)").Scan(&totalFines)

	// Available books = sum of all book copies - borrowed books
	var availableBooks int64
	r.db.Model(&domain.Book{}).Select("SUM(total_copies)").Scan(&availableBooks)
	availableBooks = availableBooks - borrowedBooks

	stats = domain.InventoryStats{
		TotalBooks:      totalBooks,
		AvailableBooks:  availableBooks,
		BorrowedBooks:   borrowedBooks,
		OverdueBooks:    overdueBooks,
		TotalStudents:   totalStudents,
		ActiveStudents:  activeStudents,
		PendingRequests: pendingRequests,
		TotalFines:      totalFines,
	}

	return &stats, nil
}
