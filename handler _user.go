package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pravytecy/rssaggregator/internal/db"
)

func(apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type params struct{
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	pa := params{}
	err:=decoder.Decode(&pa)
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	user,err := apiCfg.DB.CreateUser(r.Context(),db.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: pa.Name,
	})
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Couldn't create User: %v", err))
		return
	}
	respondWithJson(w,201,user)
}

func(apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request,user db.Customer){
	respondWithJson(w,200,user)
}