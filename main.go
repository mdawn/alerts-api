package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var pizza = []string{"Enjoy your pizza with some delicious"}

func info() {
	app.Name = "Alerting Tool"
	app.Usage = "Runs checks on API"
	app.Author = "mdawn"
	app.Version = "1.0.0"
}

var app = cli.NewApp()

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type Ticker struct {
	Open 	string `json:"open"`
	High	string `json:"high"`
	Low string `json:"low"`
	Close string `json:"close"`
	Changes []string `json:"changes"`
}

type Pricefeed []struct {
	Pair  string   `json:"pair"`
	Price 	string `json:"price"`
	PercentChange24h	string `json:"percentChange24h"`
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "btcusd",
			Aliases: []string{"b"},
			Usage:   "Ticker for btcusd",
			Action: func(c *cli.Context) {
				resp, err := http.Get("https://api.gemini.com/v2/ticker/btcusd")
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				}
			
				responseData, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
			
				var responseObject Ticker
				json.Unmarshal(responseData, &responseObject)
			
				currentPriceFake := 4000
			
				fmt.Println("Open price: ", responseObject.Open)
				fmt.Println("High price: ", responseObject.High)
				fmt.Println("Low price: ", responseObject.Low)
				fmt.Println("Current price: ", currentPriceFake)
				fmt.Println("Hourly prices per last 24 hours :", responseObject.Changes)
			},
		},
		{
			Name:    "currentPrice",
			Aliases: []string{"c"},
			Usage:   "Get current price for btcusd pair",
			Action: func(c *cli.Context) {
				resp, err := http.Get("https://api.gemini.com/v1/pricefeed")

				responseData, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
			
				var responseObject Pricefeed
				json.Unmarshal(responseData, &responseObject)
				fmt.Println(responseObject)
			},
		},
	}
}


