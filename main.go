package main

import (
	"log"
  "os"
  
  "github.com/urfave/cli"
)

func info() {
  app.Name = "Gemini API CLI"
  app.Usage = "An example CLI for some Gemini public APIs"
  app.Author = "mdawn" 
  app.Version = "1.0.0"
}


var app = cli.NewApp()
func main() {
  info()
  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
