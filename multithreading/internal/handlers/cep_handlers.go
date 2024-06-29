package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/sandersoncoelho/go-expert/multithreading/internal/entity"
)

type CepHandler struct {}

func NewCepHandler() *CepHandler {
	return &CepHandler{}
}

func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, "GET", "https://brasilapi.com.br/api/cep/v1/" + cep, nil)
	handleError(err, nil)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	msg := "Erro na requesição para economia.awesomeapi.com.br"
	handleError(err, &msg)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handleError(err, nil)

	var cepInput entity.CepBrasilApi
	err = json.Unmarshal(body, &cepInput)
	handleError(err, nil)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cepInput)
}

func handleError(err error, msg *string) {
	if err != nil {
		if msg != nil {
			println(*msg)
		}
		panic(*msg)
	}
}