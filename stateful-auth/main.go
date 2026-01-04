package main

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Username != "admin" || req.Password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	sessionID := createSession(req.Username)

	http.SetCookie(w, &http.Cookie{
		Name:     "SESSION_ID",
		Value:    sessionID,
		HttpOnly: true,
		Path:     "/",
	})

	w.Write([]byte("Logged in successfully"))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are authenticated via session"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("SESSION_ID")
	if err == nil {
		deleteSession(cookie.Value)
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "SESSION_ID",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	w.Write([]byte("Logged out"))
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.Handle("/protected", sessionMiddleware(http.HandlerFunc(protectedHandler)))
	http.HandleFunc("/logout", logoutHandler)

	http.ListenAndServe(":8080", nil)
}
