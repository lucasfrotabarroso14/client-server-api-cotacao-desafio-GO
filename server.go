// #cgo CFLAGS: -g

package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ExchangeRate struct {
	USDBRL CurrencyInfo `json:"USDBRL"`
}

type CurrencyInfo struct {
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
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dollar_quotes (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("iniciando...")
	dollarPrice, err := GetDollarPrice()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = saveToDatabase(dollarPrice["current_dollar"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	jsounOutput, err := json.Marshal(dollarPrice)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsounOutput)
	fmt.Println("Request finalizada")

}

func GetDollarPrice() (map[string]string, error) { //nessa funcao quero apenas retornar o bid
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3000*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	var exchangeRate ExchangeRate

	err = json.Unmarshal(body, &exchangeRate)

	if err != nil {
		return nil, err
	}
	currentDollar := map[string]string{
		"current_dollar": exchangeRate.USDBRL.Bid,
	}
	return currentDollar, nil

}
func saveToDatabase(bid string) error {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3000*time.Millisecond)
	defer cancel()
	_, err := db.ExecContext(ctx, "INSERT INTO dollar_quotes (bid) values (?)", bid)
	if err != nil {
		return err
	}
	fmt.Println("Dados inseridos com sucesso no banco")
	return nil

}
