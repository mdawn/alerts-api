package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
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

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "peppers",
			Aliases: []string{"p"},
			Usage:   "Add peppers to your pizza",
			Action: func(c *cli.Context) {
				pe := "peppers"
				peppers := append(pizza, pe)
				m := strings.Join(peppers, " ")
				fmt.Println(m)
			},
		},
		{
			Name:    "pineapple",
			Aliases: []string{"pa"},
			Usage:   "Add pineapple to your pizza",
			Action: func(c *cli.Context) {
				pa := "pineapple"
				pineapple := append(pizza, pa)
				m := strings.Join(pineapple, " ")
				fmt.Println(m)
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


