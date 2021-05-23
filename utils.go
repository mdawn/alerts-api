package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Symbol  string   `json:"symbol"`
	Open 	string `json:"open"`
	High	string `json:"high"`
	Low string `json:"low"`
	Close string `json:"close"`
	Changes []string `json:"changes"`
}

func btcusd() {
	resp, err := http.Get("https://api.gemini.com/v2/ticker/btcusd")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	currentPriceFake := 4000

	fmt.Println(responseObject.Symbol)
	fmt.Println(responseObject.High)
	fmt.Println(responseObject.Low)
	fmt.Printf("%s\n", responseObject.Changes)
	fmt.Println(currentPriceFake)

}

type Pricefeed struct {
	Pair  string   `json:"pair"`
	Price 	string `json:"price"`
	PercentChange24h	string `json:"percentChange24h"`
}

func currentPrice() {
	resp, err := http.Get("https://api.gemini.com/v1/pricefeed")

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Pricefeed 
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Price)

}
