package main

import (
	"fmt"
	"net/http"

	"github.com/pravytecy/rssaggregator/internal/auth"
	"github.com/pravytecy/rssaggregator/internal/db"
)

type authHandler func(w http.ResponseWriter,r *http.Request,user db.Customer)

func ( cfg*apiConfig)middleWareAuth(handler authHandler) http.HandlerFunc {
	return func (w http.ResponseWriter,r *http.Request)  {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w,403,fmt.Sprintf("Auth error: %v", err))
			return
		}
		user, err := cfg.DB.GetCustomers(r.Context(),apiKey)
		if err != nil {
			respondWithError(w,400,fmt.Sprintf("Couldn't get user: %v", err))
			return
		}
		handler(w,r,user)
	}
}