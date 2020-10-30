package db

import "gorm.io/gorm"

// Expense is a single expense, associated with a category.
type Expense struct {
	gorm.Model
	MonthYear
	Name       string
	Amount     int
	CategoryID int
	Category   *Category `gorm:"foreignKey:CategoryID"`
}
