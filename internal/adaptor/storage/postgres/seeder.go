package postgres

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
	util "github.com/sugaml/lms-api/internal/core/utils"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) {
	// -------------------------
	// 1. Seed Roles
	// -------------------------
	roles := []domain.Role{
		{Name: "STUDENT"},
		{Name: "TEACHER"},
		{Name: "ADMIN"},
		{Name: "DIRECTOR"},
		{Name: "LIBRARIAN"},
		{Name: "STAFF"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, domain.Role{Name: role.Name}).Error; err != nil {
			logrus.Error("Failed to seed role:", role.Name, err)
		}
	}
	logrus.Info("Roles seeded successfully")

	// -------------------------
	// 2. Seed Users
	// -------------------------
	users := []domain.User{
		{
			Username: "administrator",
			Password: "administrator",
			Email:    "administrator@gmail.com",
			FullName: "Administrator",
			IsActive: true,
		},
		{
			Username: "librarian",
			Password: "librarian",
			Email:    "librarian@gmail.com",
			FullName: "Librarian",
			IsActive: true,
		},
	}

	for i := range users {
		pwd, err := util.HashPassword(users[i].Password)
		if err != nil {
			logrus.Error("Password hash failed:", err)
			continue
		}
		users[i].Password = pwd

		if err := db.FirstOrCreate(
			&users[i],
			domain.User{Username: users[i].Username},
		).Error; err != nil {
			logrus.Error("Failed to seed user:", users[i].Username, err)
		}
	}
	logrus.Info("Users seeded successfully")

	// -------------------------
	// 3. Assign Roles to Users
	// -------------------------

	var adminUser domain.User
	var adminRole domain.Role

	db.First(&adminUser, "username = ?", "administrator")
	db.First(&adminRole, "name = ?", "ADMIN")
	logrus.Infof("Admin User ID: %s", adminUser.ID)
	// administrator → ADMIN
	if err := db.Model(&adminUser).Association("Roles").Append(&adminRole); err != nil {
		logrus.Error("Failed to assign ADMIN role:", err)
	}

	var librarianUser domain.User
	var librarianRole domain.Role

	db.First(&librarianUser, "username = ?", "librarian")
	db.First(&librarianRole, "name = ?", "LIBRARIAN")

	// librarian → LIBRARIAN
	if err := db.Model(&librarianUser).Association("Roles").Append(&librarianRole); err != nil {
		logrus.Error("Failed to assign LIBRARIAN role:", err)
	}

	logrus.Info("User roles seeded successfully")
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
