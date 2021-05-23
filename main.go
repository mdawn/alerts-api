package main

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli"
	"github.com/aclements/go-moremath/stats"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
  
  }

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
				log.Println("Getting BTCUSD Ticker API Data")
				resp, err := http.Get("https://api.gemini.com/v2/ticker/btcusd")
				if err != nil {
					log.Error(err)
				}

				responseData, err := ioutil.ReadAll(resp.Body)
				log.Println("Reading Response")
				if err != nil {
					log.Fatal(err)
				}

				var responseObject Ticker
				json.Unmarshal(responseData, &responseObject)
				log.Println("Parsing Response")

				log.Println("Open price: ", responseObject.Open)
				log.Println("High price: ", responseObject.High)
				log.Println("Low price: ", responseObject.Low)
				log.Println("Hourly prices per last 24 hours :", responseObject.Changes)
			},
		},
		{
			Name:    "currentPrice",
			Aliases: []string{"c"},
			Usage:   "Get current price for BTCUSD pair",
			Action: func(c *cli.Context) {
				log.Println("Getting Pricefeed API Data")
				resp, err := http.Get("https://api.gemini.com/v1/pricefeed")

				responseData, err := ioutil.ReadAll(resp.Body)
				log.Println("Reading Response")
				if err != nil {
					log.Fatal(err)
				}

				currentPriceFake := 4000.00
				log.Println("Current price: ", currentPriceFake)

				var responseObject Pricefeed
				json.Unmarshal(responseData, &responseObject)
				log.Println("Parsing Response")
				log.Println(responseObject)
			},
		},
		{
			Name:    "average",
			Aliases: []string{"v"},
			Usage:   "Get the average price for BTCUSD pair over the last 24 hours",
			Action: func(c *cli.Context) {
				log.Println("Getting BTCUSD Ticker API Data")
				resp, err := http.Get("https://api.gemini.com/v2/ticker/btcusd")
				if err != nil {
					log.Error(err)
				}

				responseData, err := ioutil.ReadAll(resp.Body)
				log.Println("Reading Response")
				if err != nil {
					log.Fatal(err)
				}

				var responseObject Ticker
				json.Unmarshal(responseData, &responseObject)
				log.Println("Parsing Response")
				hourlies := ConvertSlice(responseObject.Changes)
				log.Println("Converting Data")
				average := Average(hourlies)
				log.Println("Calculating Average")

				log.Println("Average: ", average)
			},
		},
		{
			Name:    "deviation",
			Aliases: []string{"d"},
			Usage:   "Get standard deviation for BTCUSD pair",
			Action: func(c *cli.Context) {
				log.Println("Getting BTCUSD Ticker API Data")
				resp, err := http.Get("https://api.gemini.com/v2/ticker/btcusd")
				if err != nil {
					log.Error(err)
				}

				responseData, err := ioutil.ReadAll(resp.Body)
				log.Println("Reading Response")
				if err != nil {
					log.Fatal(err)
				}

				currentPriceFake := 4000.00
				log.Println("Current price: ", currentPriceFake)

				var responseObject Ticker
				json.Unmarshal(responseData, &responseObject)
				log.Println("Parsing Response")
				hourlies := ConvertSlice(responseObject.Changes)
				log.Println("Converting Data")
				deviation := StdDev(hourlies)
				log.Println("Calculating Average")
				average := Average(hourlies)
				
				log.Println("Average: ", average)

				log.Println("Standard deviation: ", deviation)

				log.Println("Calculating deviation from Average")
				if currentPriceFake > (average + deviation) {
					log.Println("Current Price is > 1 StdDev from Avg")
				}
				if currentPriceFake < (average - deviation) {
					log.Println("Current Price is < 1 StdDev from Avg")
				}
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
			log.Fatal(err)
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
