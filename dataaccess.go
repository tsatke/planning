package planning

import (
	"fmt"

	"planning/db"
)

// dataAccess acts as a dataAccess between server and database.
// Data requests from the frontend pass through this dataAccess
// before they hit the actual database.
type dataAccess struct {
	db *db.DB
}

func (da dataAccess) GetCategory(name string) *db.Category {
	return da.db.FindCategory(name)
}

func (da dataAccess) CreateCategory(name string) (*db.Category, error) {
	if da.db.HasCategory(name) {
		return nil, fmt.Errorf("category already exists")
	}
	return da.db.CreateCategory(name)
}

func (da dataAccess) AllCategories() []*db.Category {
	return da.db.AllCategories()
}

func (da dataAccess) DeleteCategory(name string) error {
	if !da.db.HasCategory(name) {
		return fmt.Errorf("category doesn't exist")
	}
	if !da.db.DeleteCategory(name) {
		return fmt.Errorf("unable to delete category")
	}
	return nil
}

func (da dataAccess) RenameCategory(old, new string) error {
	if err := da.db.RenameCategory(old, new); err != nil {
		if err == db.ErrDoesntExist {
			return fmt.Errorf("old category doesn't exist")
		} else if err == db.ErrAlreadyExists {
			return fmt.Errorf("new category already exists")
		}
		return fmt.Errorf("unknokwn error while renaming: %w", err)
	}
	return nil
}
