package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Cotacao struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type CotacaoMux struct {
	DB *sql.DB
}

func handleError(err error, msg *string) {
	if err != nil {
		if msg != nil {
			println(*msg)
		}
		panic(*msg)
	}
}

func getCotacao() Cotacao {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 2000)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	handleError(err, nil)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	msg := "Erro na requesição para economia.awesomeapi.com.br"
	handleError(err, &msg)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handleError(err, nil)

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	handleError(err, nil)
	return cotacao
}

func (cotacaoMux CotacaoMux) persistCotacao(cotacao Cotacao) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond * 10)
	defer cancel()

	tx, err := cotacaoMux.DB.BeginTx(ctx, nil)
	handleError(err, nil)

	stmt, err := tx.Prepare("insert into COTACAO(BID) values(?)")
	msg := "Erro ao persistir no banco de dados"
	handleError(err, &msg)
	defer stmt.Close()

	_, err = stmt.Exec(cotacao.USDBRL.Bid)
	handleError(err, nil)

	err = tx.Commit()
	handleError(err, nil)
}

func (cotacaoMux CotacaoMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cotacao := getCotacao()

	cotacaoMux.persistCotacao(cotacao)
	
	fmt.Fprintf(w, "{\"bid\": %s}", cotacao.USDBRL.Bid)
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, r.(string), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func main()  {
	db, err := sql.Open("sqlite3", "./cotacao.db")
	msg := "Falha ao conectar com o banco de dados"
	handleError(err, &msg)
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("/cotacao", CotacaoMux{DB: db})

	log.Fatal(http.ListenAndServe(":8080", recoverMiddleware(mux)))
}