package postgres

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
	util "github.com/sugaml/lms-api/internal/core/utils"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	users := []domain.User{
		{
			Username: "administrator",
			Password: "administrator",
			Email:    "administrator@gmail.com",
			FullName: "Administrator",
			Role:     "administrator",
			IsActive: true,
		},
		{
			Username: "librarian",
			Email:    "librarian@gmail.com",
			FullName: "Librarian",
			Password: "librarian",
			Role:     "librarian",
			IsActive: true,
		},
	}
	for _, user := range users {
		pwd, err := util.HashPassword(user.Password)
		if err != nil {
			logrus.Error("Failed to seed user:", user, err)
		}
		user.Password = pwd
		if err := db.FirstOrCreate(&user, domain.User{Username: user.Username, Email: user.Email, Password: user.Password, Role: user.Role}).Error; err != nil {
			logrus.Error("Failed to seed user:", user, err)
		}
	}
	logrus.Info("Users seeded successfully")
}

func SeedCategories(db *gorm.DB) {
	categories := []string{
		"Finance", "Marketing", "Management", "Economics",
		"Accounting", "Operations", "Human Resources",
		"Information Technology", "Communication",
		"Information Systems", "Sales", "Statistics",
	}

	for _, name := range categories {
		category := &domain.Category{Name: name, Slug: GenerateSlug(name), IsActive: true}
		if err := db.FirstOrCreate(category, domain.Category{Name: name, Slug: GenerateSlug(name)}).Error; err != nil {
			logrus.Error("Failed to seed category:", name, err)
		}
	}
	logrus.Info("Categories seeded successfully")
}

func SeedPrograms(db *gorm.DB) {
	programs := []string{
		"MBA", "MBA IT", "MBA Finance", "MBA GLM",
	}
	for _, name := range programs {
		program := &domain.Program{Name: name, Slug: GenerateSlug(name), IsActive: true}
		if err := db.FirstOrCreate(program, domain.Program{Name: name, Slug: GenerateSlug(name)}).Error; err != nil {
			logrus.Error("Failed to seed program:", name, err)
		}
	}
	logrus.Info("Programs seeded successfully")
}

func GenerateSlug(name string) string {
	// Trim leading/trailing spaces
	slug := strings.TrimSpace(name)
	// Convert to lowercase
	slug = strings.ToLower(slug)
	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
