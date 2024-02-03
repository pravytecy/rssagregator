package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main()  {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port not found")
	}
	fmt.Println("Port : ",port)

	router := chi.NewRouter();
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	v1router :=  chi.NewRouter();
	v1router.Get("/healthz",handlerReadiness)
	v1router.Get("/error",handlerError)
	router.Mount("/v1",v1router)
	serve := &http.Server{
		Handler: router,
		Addr: ":"+port,
	}
	log.Printf("Server is running  on port %v",port)
	err := serve.ListenAndServe();
	if err != nil {
		log.Fatal(err)
	}
}