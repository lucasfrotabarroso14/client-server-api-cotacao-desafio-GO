package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://127.0.0.1:8080/", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			log.Fatalf("Erro ao fazer a requisicao GET para o server: %v", err)
		}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	currentDollarStr, ok := result["current_dollar"].(string)
	if !ok {
		log.Fatal("erro")
	}
	currentDollar, err := strconv.ParseFloat(currentDollarStr, 64)
	if err != nil {
		log.Fatal("erro na conversao para float")

	}
	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	output := fmt.Sprintf("Dolar: %.2f", currentDollar)
	_, err = file.Write([]byte(output))
	if err != nil {
		panic(err)
	}
	fmt.Println("Valor de current_dollar escrito no arquivo 'cotacao.txt'")

}
