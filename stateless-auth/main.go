package main

import (
	"encoding/json"
	"fmt"
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

	token, _ := generateToken(req.Username)

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You have accessed a protected route"))
}

func main() {
	fmt.Println("Starting server at port 8080")
	http.HandleFunc("/login", loginHandler)
	http.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}
