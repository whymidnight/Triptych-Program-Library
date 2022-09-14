package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	giveawaysActions "triptych.labs/giveaways/v2/actions"
	giveawaysDatabase "triptych.labs/giveaways/v2/database"
	giveawaysJobs "triptych.labs/giveaways/v2/jobs"
	giveawaysState "triptych.labs/giveaways/v2/state"

	twitterActions "triptych.labs/twitter/v2/actions"
	twitterConstants "triptych.labs/twitter/v2/constants"
	twitterState "triptych.labs/twitter/v2/state"

	"triptych.labs/utils"
)

func init() {
	twitterConstants.Init()
	twitterState.Init()
	giveawaysDatabase.Init()
  giveawaysJobs.Init()
  giveawaysState.Init()
}

func main() {
	go func() {
		var exitCh = make(chan os.Signal)
		signal.Notify(exitCh, syscall.SIGINT)
		signal.Notify(exitCh, syscall.SIGTERM)
		s := <-exitCh
		fmt.Printf("Caught signal %v. Exiting ...\n", s)

		giveawaysJobs.Stop()
		os.Exit(0)
	}()

	fmt.Println("RPC:", utils.NETWORK)

	if twitterConstants.ConsumerKey == "" && twitterConstants.ConsumerSecret == "" {
		fmt.Println("Please setup ConsumerKey and ConsumerSecret.")
		return
	}

	var port *int = flag.Int(
		"port",
		8888,
		"Port to listen on.")
	flag.Parse()

	http.HandleFunc("/twitter", twitterActions.Index)
	http.HandleFunc("/twitter/maketoken", twitterActions.GetTwitterToken)
	http.HandleFunc("/twitter/request", twitterActions.RedirectUserToTwitter)
	http.HandleFunc("/twitter/verify", twitterActions.Verify)

	http.HandleFunc("/giveaways", giveawaysActions.Index)

	u := fmt.Sprintf(":%d", *port)
	fmt.Printf("Listening on '%s'\n", u)
	http.ListenAndServe(u, nil)


}
