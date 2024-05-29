package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Cotacao struct {
	Bid float64 `json:"bid"`
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writeData(cotacao Cotacao) {
	file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	handleError(err)
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("Dólar:%f\n", cotacao.Bid))
	handleError(err)
 }

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 3000)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	handleError(err)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	handleError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handleError(err)

	if resp.StatusCode == http.StatusInternalServerError {
		log.Fatal(string(body))
	}

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	handleError(err)

	writeData(cotacao)
}