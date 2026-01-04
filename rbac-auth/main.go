package main

import "net/http"

func adminHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Admin access granted"))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User access granted"))
}

func main() {
	admin := authorize(RoleAdmin)(http.HandlerFunc(adminHandler))
	user := authorize(RoleUser)(http.HandlerFunc(userHandler))

	http.Handle("/admin", authMiddleware(admin))
	http.Handle("/user", authMiddleware(user))

	http.ListenAndServe(":8080", nil)
}
