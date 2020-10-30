package planning

import (
	"fmt"

	"planning/db"
	"planning/server"

	"github.com/rs/zerolog"
)

type Application struct {
	log zerolog.Logger
	db  *db.DB
}

func New(opts ...Opt) *Application {
	app := &Application{}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

func (app Application) Run() error {
	db, err := db.Open(app.log, "test.db")
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer func() { _ = db.Close() }()

	srv := server.New(app.log, ":51560", dataAccess{db})
	if err := srv.Start(); err != nil {
		return err
	}

	return nil
}
