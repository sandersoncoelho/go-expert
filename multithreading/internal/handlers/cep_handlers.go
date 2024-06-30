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

func getCepBrasilApi(cep string) entity.CepBrasilApi {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, "GET", "https://brasilapi.com.br/api/cep/v1/" + cep, nil)
	handleError(err, nil)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	msg := "Erro na requesição para brasilapi"
	handleError(err, &msg)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handleError(err, nil)

	var cepResponse entity.CepBrasilApi
	err = json.Unmarshal(body, &cepResponse)
	handleError(err, nil)

	return cepResponse
}

func getCepViaCep(cep string) entity.CepViaCep {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, "GET", "https://viacep.com.br/ws/" + cep + "/json/", nil)
	handleError(err, nil)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	msg := "Erro na requesição para viacep"
	handleError(err, &msg)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handleError(err, nil)

	var cepResponse entity.CepViaCep
	err = json.Unmarshal(body, &cepResponse)
	handleError(err, nil)

	return cepResponse
}

func (h *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c1 := make(chan any)
	c2 := make(chan any)

	go func() {
		c1 <- getCepBrasilApi(cep)
	}()

	go func() {
		c2 <- getCepViaCep(cep)
	}()

	select {
	case cepResponse := <- c1:
		println("brasilapi")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(entity.CepResponse{ServiceName: "brasilapi", Data: cepResponse})

	case cepResponse := <- c2:
		println("viacep")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(entity.CepResponse{ServiceName: "viacep", Data: cepResponse})

	case <- time.After(time.Second):
		println("timeout")
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Request timeout\n"))
	}
}

func handleError(err error, msg *string) {
	if err != nil {
		if msg != nil {
			println(*msg)
		}
		panic(*msg)
	}
}