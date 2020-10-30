package db

import (
	"gorm.io/gorm"
)

// Category is a category, which expenses or budgets can be associated
// with.
type Category struct {
	gorm.Model
	Name string
}

func (db *DB) CreateCategory(name string) (*Category, error) {
	if db.FindCategory(name) != nil {
		return nil, ErrAlreadyExists
	}
	created := &Category{
		Name: name,
	}
	db.store.Create(created)
	return created, nil
}

func (db *DB) FindCategory(name string) *Category {
	var category Category
	tx := db.store.Find(&category, "name = ?", name)
	if tx.RowsAffected == 0 {
		return nil
	}
	return &category
}

func (db *DB) HasCategory(name string) bool {
	return db.FindCategory(name) != nil
}

func (db *DB) AllCategories() []*Category {
	var categories []*Category
	db.store.Find(&categories)
	return categories
}

func (db *DB) DeleteCategory(name string) bool {
	category := db.FindCategory(name)
	if category == nil {
		return false
	}
	return db.store.Delete(&category).RowsAffected > 0
}

func (db *DB) RenameCategory(old, new string) error {
	category := db.FindCategory(old)
	if category == nil {
		return ErrDoesntExist
	}
	if db.HasCategory(new) {
		return ErrAlreadyExists
	}
	category.Name = new
	db.store.Save(category)
	return nil
}
