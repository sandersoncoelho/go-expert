package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sandersoncoelho/go-expert/multithreading/internal/handlers"
)

func main() {
	cepHandler := handlers.NewCepHandler()

	r := chi.NewRouter()
	r.Get("/ceps", cepHandler.GetCep)
	
	http.ListenAndServe(":8000", r)
}