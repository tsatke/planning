package server

import "planning/db"

type Repository interface {
	CreateCategory(name string) (*db.Category, error)
	GetCategory(name string) *db.Category
	AllCategories() []*db.Category
	DeleteCategory(name string) error
	RenameCategory(old, new string) error
}
