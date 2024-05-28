package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

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

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (cotacaoMux CotacaoMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	handleError(err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	handleError(err)

	var cotacao Cotacao
	err = json.Unmarshal(body, &cotacao)
	handleError(err)


	tx, err := cotacaoMux.DB.Begin()
	handleError(err)

	stmt, err := tx.Prepare("insert into COTACAO(BID) values(?)")
	handleError(err)
	defer stmt.Close()

	_, err = stmt.Exec(cotacao.USDBRL.Bid)
	handleError(err)

	err = tx.Commit()
	handleError(err)
	

	fmt.Fprintf(w, "{\"bid\": %s}", cotacao.USDBRL.Bid)
}

func main()  {
	db, err := sql.Open("sqlite3", "./cotacao.db")
	handleError(err)
	defer db.Close()

	mux := http.NewServeMux()
	mux.Handle("/cotacao", CotacaoMux{DB: db})

	log.Fatal(http.ListenAndServe(":8080", mux))
}