package main

import (
	"crypto/rand"
	"encoding/hex"
)

var sessions = make(map[string]string)

func createSession(email string) string {
	b := make([]byte, 32)
	rand.Read(b)
	id := hex.EncodeToString(b)
	sessions[id] = email
	return id
}

func getSession(id string) (string, bool) {
	email, ok := sessions[id]
	return email, ok
}
