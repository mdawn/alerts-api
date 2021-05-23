To run this code:

- use a Linux-based machine (Mac plz)
- run `go version`. If the result is not `go version go1.16.3 darwin/amd64`, go ahead and get that sorted out.
- clone this repo
- run `go get` and `go mod tidy`
- `go build main.go`
- run `cp main /usr/local/bin`
-  then you can run commands like `main — help`

Returns

```shell
NAME:
   Gemini API CLI - An example CLI for some Gemini public APIs

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   mdawn

COMMANDS:
   peppers, p     Add peppers to your pizza
   pineapple, pa  Add pineapple to your pizza
   cheese, c      Add cheese to your pizza
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

And something like

`main p`

Returns

`Enjoy your pizza with some delicious peppers`