package planning

import "github.com/rs/zerolog"

type Opt func(*Application)

func WithLogger(log zerolog.Logger) Opt {
	return func(app *Application) {
		app.log = log
	}
}
