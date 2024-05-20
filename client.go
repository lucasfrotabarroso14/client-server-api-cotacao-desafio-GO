package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	resp, err := http.Get("http://127.0.0.1:8080/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// abaixo eu estou declarando um map com a chave string e o valor de qualquer tipp
	// em GO uma interface vazia  indica que os valores podem ser de qualquer tipo
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
