package main

import (
	"encoding/json"
	"fmt"

	"io/ioutil"

	"net/http"
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

func main() {
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

	resp, error := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if error != nil {
		return nil, error
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
