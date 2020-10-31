package db

import (
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	log   zerolog.Logger
	store *gorm.DB
}

func Open(log zerolog.Logger, dbFile string) (*DB, error) {
	dialector := sqlite.Open(dbFile)
	log.Debug().
		Str("dialect", dialector.Name()).
		Str("file", dbFile).
		Msg("open database")
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("gorm open: %w", err)
	}

	log.Debug().
		Msg("auto migrate schema")
	if err := db.AutoMigrate(
		&Expense{},
		&Category{},
		&Budget{},
	); err != nil {
		return nil, fmt.Errorf("auto migrate: %w", err)
	}
	return &DB{
		log:   log,
		store: db,
	}, nil
}

func (db DB) Close() error {
	return nil
}
