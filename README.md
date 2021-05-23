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
   Alerting Tool - Runs checks on API

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   mdawn

COMMANDS:
   btcusd, b        Ticker for btcusd
   currentPrice, c  Get current price for btcusd pair
   help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

And something like

`main b`

Returns

```shell
Open price:  36410.02
High price:  38861.15
Low price:  35272.09
Current price:  4000
Hourly prices per last 24 hours : [37484.18 37491.88 37933.64 38214.35 37496.25 37669 38230.38 38087.2 37371.75 37668.33 38026.06 38189.33 38205.1 37837.11 38339.84 38204.88 38175.16 38387.95 37220.67 36885.24 36333.93 36452.26 36434.99 36356.37]
```