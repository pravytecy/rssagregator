package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string){
	if statusCode > 499 {
		log.Printf("Something 5XX is occured %v",msg)
	}
	type errorResponse struct{
		Error string `json:"error"`
	}
	respondWithJson(w,statusCode,errorResponse{Error: msg})
}

func respondWithJson(w http.ResponseWriter, statusCode int, payload interface{}){
	data,err := json.Marshal(payload)
	log.Printf("Failed to marshall response %v",payload)
	if err != nil{
		w.WriteHeader(500)
		return
	}
	w.Header().Add("content-type","application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}