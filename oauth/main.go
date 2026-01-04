package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/oauth2"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := exchangeCode(code)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	client := oauthConfig.Client(r.Context(), token)
	resp, _ := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	defer resp.Body.Close()

	var user map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&user)

	email := user["email"].(string)
	sessionID := createSession(email)

	http.SetCookie(w, &http.Cookie{
		Name:     "SESSION_ID",
		Value:    sessionID,
		HttpOnly: true,
		Path:     "/",
	})

	w.Write([]byte("OAuth login successful"))
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Authenticated via OAuth2"))
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/oauth/callback", callbackHandler)
	http.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))

	http.ListenAndServe(":8080", nil)
}
