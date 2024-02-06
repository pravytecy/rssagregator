package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pravytecy/rssaggregator/internal/db"
)

func(apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user db.Customer){
	type params struct{
		Name string `json:"name"`
		Url string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	pa := params{}
	err:=decoder.Decode(&pa)
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feed,err := apiCfg.DB.CreateFeed(r.Context(),db.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: pa.Name,
		Url: pa.Url,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Couldn't create Feed: %v", err))
		return
	}
	respondWithJson(w,201,feed)
}


func(apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request){
	feeds,err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	// feeds := Fee
	// for _,feed := range dbFeeds {
		
	// }
	respondWithJson(w,201,feeds)
}
