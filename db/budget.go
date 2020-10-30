package db

import "gorm.io/gorm"

// Budget is a budget, i.e. limitation (soft, more as a hint) on a
// certain category.
type Budget struct {
	gorm.Model
	Amount     int
	CategoryID int
	Category   *Category `gorm:"foreignKey:CategoryID"`
}
