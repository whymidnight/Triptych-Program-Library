package main

import (
	"flag"
	"fmt"
	"net/http"

	"triptych.labs/twitter/v2/constants"
	"triptych.labs/twitter/v2/state"
	"triptych.labs/utils"
)

func init() {
	constants.Init()
	state.Init()
}

func main() {

	fmt.Println("RPC:", utils.NETWORK)

	if constants.ConsumerKey == "" && constants.ConsumerSecret == "" {
		fmt.Println("Please setup ConsumerKey and ConsumerSecret.")
		return
	}

	var port *int = flag.Int(
		"port",
		8888,
		"Port to listen on.")

	flag.Parse()

	u := fmt.Sprintf(":%d", *port)
	fmt.Printf("Listening on '%s'\n", u)
	http.ListenAndServe(u, nil)
}


