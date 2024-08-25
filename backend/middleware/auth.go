package middleware

import (
	"net/http"

	"github.com/aeilang/backend/config"
)

func Auth(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := config.GetConfig()
		auth := r.Header.Get("Authorization")
		if auth != cfg.TOKEN {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
