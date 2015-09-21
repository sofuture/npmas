package main

import (
	"flag"
	"os"

	"github.com/sofuture/npmas/client"
)

var server string

func main() {
	flag.StringVar(&server, "server", "localhost:6600", "mpd server to connect to")

	flag.Parse()

	os.Exit(client.Run(server))
}
