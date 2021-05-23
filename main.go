package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

type Response struct {
	Symbol  string   `json:"symbol"`
	Open 	string `json:"open"`
	High	string `json:"high"`
	Low string `json:"low"`
	Close string `json:"close"`
	Changes []string `json:"changes"`
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
			
				var responseObject Response
				json.Unmarshal(responseData, &responseObject)
			
				currentPriceFake := 4000
			
				fmt.Println(responseObject.Symbol)
				fmt.Println(responseObject.High)
				fmt.Println(responseObject.Low)
				fmt.Printf("%s\n", responseObject.Changes)
				fmt.Println(currentPriceFake)
			},
		},
		{
			Name:    "cheese",
			Aliases: []string{"c"},
			Usage:   "Add cheese to your pizza",
			Action: func(c *cli.Context) {
				ch := "cheese"
				cheese := append(pizza, ch)
				m := strings.Join(cheese, " ")
				fmt.Println(m)
			},
		},
	}
}


