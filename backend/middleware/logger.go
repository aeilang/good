package middleware

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type responseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.StatusCode = statusCode
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		StatusCode:     http.StatusOK,
	}
}

func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		reqeustID := GetRequestID(r.Context())

		lrw := newResponseWriter(w)
		defer func() {

			statusString := http.StatusText(lrw.StatusCode)
			log.Info().
				Str("id", reqeustID).
				Str("method", r.Method).
				Str("path", r.URL.Path).
				Int("status_code", lrw.StatusCode).
				Dur("elapsed_ms", time.Since(start)).
				Msg(statusString)
		}()

		next.ServeHTTP(lrw, r)
	}

	return http.HandlerFunc(fn)
}
