package postgres

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/adaptor/config"
	"github.com/sugaml/lms-api/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database connection
func NewDB(config config.Config) (*gorm.DB, error) {
	newLogger := logger.New(
		logrus.New(),
		logger.Config{
			SlowThreshold: time.Second * 2,
			LogLevel:      logger.Warn,
			Colorful:      true,
		},
	)
	db, err := gorm.Open(postgres.Open(config.DB_SOURCE), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		DryRun:                                   false,
		Logger:                                   newLogger,
	})
	if err != nil {
		return nil, err
	}
	if config.DB_DEBUG == "true" {
		db = db.Debug()
	}
	var dbName string
	err = db.Raw("SELECT current_database();").Scan(&dbName).Error
	if err != nil {
		return nil, err
	}
	if config.DB_AUTO_MIGRATE != "false" {
		err = db.AutoMigrate(
			&domain.User{},
			&domain.AuditLog{},
			&domain.Book{},
			&domain.BookCopy{},
			&domain.Fine{},
			&domain.BorrowedBook{},
			&domain.Category{},
			&domain.Program{},
			&domain.Notification{},
		)
		if err != nil {
			return nil, err
		}
		// db.Raw("CREATE EXTENSION IF NOT EXISTS pg_trgm;")
	}
	db.Migrator().CreateConstraint(&domain.ClassRoutine{}, "unique_room_time")
	db.Exec(`
		CREATE UNIQUE INDEX idx_room_time
		ON class_routines (room_id, day_of_week, time_slot_id);
		`)

	db.Exec(`
		CREATE UNIQUE INDEX idx_teacher_time
		ON class_routines (teacher_id, day_of_week, time_slot_id);
		`)

	db.Exec(`
		CREATE UNIQUE INDEX idx_semester_time
		ON class_routines (semester_id, day_of_week, time_slot_id);
		`)

	// Seed initial data
	SeedUsers(db)
	SeedCategories(db)
	SeedPrograms(db)
	logrus.Infof("Successfully connected to the database :: %s", dbName)
	return db, nil
}
