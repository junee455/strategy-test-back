package routes

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func TestLoginListener(w http.ResponseWriter, r *http.Request) {
	// r.Body.
	type TestLoginRequest struct {
		ClientId string `json:"clientId"`
	}

	var payload TestLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid login data", http.StatusBadRequest)
		return
	}

	var testUserData = payload

	testUserDataJson, err := json.Marshal(testUserData)

	testUserDataJson = []byte(url.QueryEscape(string(testUserDataJson)))

	if err != nil {
		panic(err)
	}

	var userCookie = &http.Cookie{
		Name:     "StAuth",
		Value:    string(testUserDataJson),
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(w, userCookie)
	w.WriteHeader(http.StatusOK)
}
