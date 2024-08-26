package main

import (
	"encoding/json"
	"net/http"
)

// what does it mean mock
// mock database to store users

var users = map[string]string{}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Fname    string `json:"fname"`
		Lname    string `json:"lname"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&user)

	users[user.Username] = user.Password

	w.WriteHeader(http.StatusCreated)
	json.NewDecoder(w).Encode(map[string]string{"message": "User created"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Username string `json:"username"`
		Password string `json:password`
	}
	json.NewDecoder(r.Body).Decode(&credentials)

	storedPassword, ok := users[credentials.Username]
	if !ok || storedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewDecoder(w).Encode(map[string]string{"message": "Invalid credentials"})
		return
	}

	//Generate
}
