package mcp

import (
	"net/http"
	"strings"
)

func (c *Controller) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if c.token == "" {
			if next != nil {
				next.ServeHTTP(w, r)
			}

			return
		}

		v := r.Header.Get("X-Hg-Token")
		if v == "" {
			v = strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		}

		if v == "" {
			w.WriteHeader(http.StatusUnauthorized)

			return
		}

		if c.token != v {
			w.WriteHeader(http.StatusForbidden)

			return
		}

		if next != nil {
			next.ServeHTTP(w, r)
		}
	})
}
