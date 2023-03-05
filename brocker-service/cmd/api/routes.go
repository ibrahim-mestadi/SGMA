package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func routes() http.Handler{
	mux := chi.NewRouter()
	
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "PUT", "DELETE", "POST"}, 
		AllowedHeaders: []string {"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}



	}))
}