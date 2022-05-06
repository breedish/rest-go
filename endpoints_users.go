package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (api apiConfig) usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		api.getUser(w, r)
	case http.MethodPost:
		api.createUser(w, r)
	case http.MethodPut:
		api.updateUser(w, r)
	case http.MethodDelete:
		api.deleteUser(w, r)
	default:
		respondWithError(w, http.StatusNotFound, errors.New("method not supported"))
	}
}

func (api apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := api.dbClient.CreateUser(params.Email, params.Password, params.Name, params.Age)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusCreated, user)
}

func (api apiConfig) getUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	if len(email) == 0 {
		respondWithError(w, http.StatusBadRequest, errors.New("no id were provided"))
		return
	}

	user, err := api.dbClient.GetUser(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (api apiConfig) updateUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	if len(email) == 0 {
		respondWithError(w, http.StatusBadRequest, errors.New("no id were provided"))
		return
	}

	type parameters struct {
		Password string `json:"password"`
		Name     string `json:"name"`
		Age      int    `json:"age"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	user, err := api.dbClient.UpdateUser(email, params.Password, params.Name, params.Age)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err)
		return
	}

	respondWithJSON(w, http.StatusOK, user)
}

func (api apiConfig) deleteUser(w http.ResponseWriter, r *http.Request) {
	email := strings.TrimPrefix(r.URL.Path, "/users/")
	if len(email) == 0 {
		respondWithError(w, http.StatusBadRequest, errors.New("no id were provided"))
		return
	}

	err := api.dbClient.DeleteUser(email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err)
		return
	}

}
