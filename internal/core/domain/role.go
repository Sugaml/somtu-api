package domain

type Role struct {
	BaseModel
	Name string `gorm:"unique;not null"`

	Users []User `gorm:"many2many:user_roles;"`
}

type UserRole struct {
	UserID string `gorm:"primaryKey"`
	RoleID string `gorm:"primaryKey"`

	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role Role `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}
