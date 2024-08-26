package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fetal("Error loading .env file")
	}

	r := mux.NewRouter()

	r.HandleFunc("/signup", SignupHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
	http.Handler('/', r)
	log.Fetal(http.ListenAndServe(":8080", nil))

}
