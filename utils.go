package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

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
