package main

import (
	"crypto/rand"
	"encoding/hex"
)

var sessions = make(map[string]string)

func createSession(username string) string {
	b := make([]byte, 32)
	rand.Read(b)
	sessionID := hex.EncodeToString(b)

	sessions[sessionID] = username
	return sessionID
}

func getUserFromSession(sessionID string) (string, bool) {
	username, ok := sessions[sessionID]
	return username, ok
}

func deleteSession(sessionID string) {
	delete(sessions, sessionID)
}
