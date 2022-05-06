package main

import (
	"encoding/json"
	"fmt"
	"github.com/breedish/http_server_golang/internal/database"
	"log"
	"net/http"
	"time"
)

type apiConfig struct {
	dbClient database.Client
}

const dbPath = "./db.json"

func main() {

	api := apiConfig{
		dbClient: database.NewClient(dbPath),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/users", api.usersHandler)
	mux.HandleFunc("/users/", api.usersHandler)

	const addr = "localhost:8080"

	server := http.Server{
		Handler:      mux,
		Addr:         addr,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Println("Server started", addr)
	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if payload != nil {
		response, err := json.Marshal(payload)

		if err != nil {
			log.Println("error marshalling", err)

			w.WriteHeader(500)
			response, _ := json.Marshal(errorBody{
				Error: "error marshalling",
			})
			w.Write(response)
			return
		}
		w.WriteHeader(code)
		w.Write(response)
	}

}

func respondWithError(w http.ResponseWriter, code int, err error) {
	if err == nil {
		log.Println("missing error")
		return
	}

	log.Println(err)
	respondWithJSON(
		w,
		code,
		errorBody{
			Error: err.Error(),
		},
	)
}

type errorBody struct {
	Error string `json:"error"`
}
