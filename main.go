package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/aclements/go-moremath/stats"
)

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
	Open    string   `json:"open"`
	High    string   `json:"high"`
	Low     string   `json:"low"`
	Close   string   `json:"close"`
	Changes []string `json:"changes"`
}

type Pricefeed []struct {
	Pair             string `json:"pair"`
	Price            string `json:"price"`
	PercentChange24h string `json:"percentChange24h"`
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "btcusd",
			Aliases: []string{"b"},
			Usage:   "Ticker for BTCUSD pair",
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

				fmt.Println("Open price: ", responseObject.Open)
				fmt.Println("High price: ", responseObject.High)
				fmt.Println("Low price: ", responseObject.Low)
				fmt.Println("Hourly prices per last 24 hours :", responseObject.Changes)
			},
		},
		{
			Name:    "currentPrice",
			Aliases: []string{"c"},
			Usage:   "Get current price for BTCUSD pair",
			Action: func(c *cli.Context) {
				resp, err := http.Get("https://api.gemini.com/v1/pricefeed")

				responseData, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				currentPriceFake := 4000
				fmt.Println("Current price: ", currentPriceFake)

				var responseObject Pricefeed
				json.Unmarshal(responseData, &responseObject)
				fmt.Println(responseObject)
			},
		},
		{
			Name:    "deviation",
			Aliases: []string{"d"},
			Usage:   "Get standard deviation for BTCUSD pair",
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
				hourlies := ConvertSlice(responseObject.Changes)
				deviation := StdDev(hourlies)

				fmt.Println("Standard deviation: ", deviation)
			},
		},
		{
			Name:    "average",
			Aliases: []string{"v"},
			Usage:   "Get the average price for BTCUSD pair over the last 24 hours",
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
				hourlies := ConvertSlice(responseObject.Changes)
				average := Average(hourlies)

				fmt.Println("Average: ", average)
			},
		},
	}
}

//
// BEGIN UTILS
//

var p []float64

// ConvertSlice converts the given slice of string values to floating points
func ConvertSlice(xs []string) (p []float64) {
	for _, arg := range xs[1:] {
		n, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			panic(err)
		}
		p = append(p, n)
	}
	return p

}

// StdDev returns the standard deviation of our floating point slice of 24 hour values
func StdDev(k []float64) float64 {
	return math.Sqrt(stats.Variance(k))
}


// Average returns the average of our floating point slice of 24 hour values
func Average(v []float64) float64 {
	var t float64 = 0
	for _, value := range v {
		t += value
	}
	return t / float64(len(v))
}
