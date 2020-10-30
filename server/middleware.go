package server

import "net/http"

func (s *Server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if pn := recover(); pn != nil {
					s.log.Error().
						Interface("panic", pn).
						Msg("recover from crash while serving request")
					w.WriteHeader(http.StatusInternalServerError)
				}
			}()
			s.log.Trace().
				Str("path", r.URL.Path).
				Str("from", r.RemoteAddr).
				Str("method", r.Method).
				Msg("request")
			next.ServeHTTP(w, r)
		},
	)
}
