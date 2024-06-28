package handlers

import (
	"net/http"
)

type CepHandler struct {}

func NewCepHandler() *CepHandler {
	return &CepHandler{}
}

func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("teste"))
}