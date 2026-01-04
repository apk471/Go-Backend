package main

import "net/http"

func authorize(required Role) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			username := r.Context().Value("username")
			if username == nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			role := getRole(username.(string))
			if role != required {
				w.WriteHeader(http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
