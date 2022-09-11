package actions

import (
	"fmt"
	"net/http"

	"triptych.labs/twitter/v2/state"
)

func RedirectUserToTwitter(w http.ResponseWriter, r *http.Request) {
	walletPublicKey := r.Header.Get("PublicKey")

	ident := state.TwitterClient[walletPublicKey]

	http.Redirect(w, r, ident.RequestUrl, http.StatusTemporaryRedirect)
	fmt.Println("Leave redirtect")
}
