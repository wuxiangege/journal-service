// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"net/http"
	"strings"
)

type CorsMiddleware struct {
	allowOrigins []string
}

func NewCorsMiddleware(allowOrigins []string) *CorsMiddleware {
	if len(allowOrigins) == 0 {
		allowOrigins = []string{"http://localhost:5173", "http://127.0.0.1:5173"}
	}
	return &CorsMiddleware{allowOrigins: allowOrigins}
}

func (m *CorsMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" && m.isAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Vary", "Origin")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func (m *CorsMiddleware) isAllowed(origin string) bool {
	for _, o := range m.allowOrigins {
		if o == "*" || strings.EqualFold(o, origin) {
			return true
		}
	}
	return false
}
