package main

import "net/http"

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("SESSION_ID")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		_, ok := getSession(cookie.Value)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
