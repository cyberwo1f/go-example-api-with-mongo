package middleware

import (
	"net/http"
)

type Middleware struct {
	cfg *Config
}

type Config struct {
	AllowedOrigins     string
	AllowedHeaders     string
	AllowedMethods     string
	AllowedCredentials string
}

func NewMiddleware(cfg *Config) *Middleware {
	return &Middleware{cfg: cfg}
}

func (m *Middleware) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", m.cfg.AllowedOrigins)
		w.Header().Set("Access-Control-Allow-Headers", m.cfg.AllowedHeaders)
		w.Header().Set("Access-Control-Allow-Methods", m.cfg.AllowedMethods)
		w.Header().Set("Access-Control-Allow-Credentials", m.cfg.AllowedCredentials)

		next.ServeHTTP(w, r)
	})
}
