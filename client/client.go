package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

type Cotacao struct {
	valor float64
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second * 5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	handleError(err)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	handleError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handleError(err)
	println(string(body))
}