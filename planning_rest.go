package planning

import (
	"fmt"

	"github.com/rs/zerolog"
	"github.com/tsatke/planning/db"
	"github.com/tsatke/planning/server"
)

type Application struct {
	log zerolog.Logger
	db  *db.DB

	openBrowser bool
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

	srv := server.New(app.log, ":0", dataAccess{db})
	if err := srv.Start(app.openBrowser); err != nil {
		return err
	}

	return nil
}
