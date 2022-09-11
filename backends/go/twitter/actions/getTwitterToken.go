package actions

import (
	"net/http"
	"time"

	"triptych.labs/twitter/v2/database"
	"triptych.labs/twitter/v2/state"
)

func GetTwitterToken(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	verificationCode := values.Get("oauth_verifier")
	tokenKey := values.Get("oauth_token")

	walletPublicKey, ok := state.TwitterAuthCodes[tokenKey]
	if ok {
		return
	}

	twitterClient := state.TwitterClient[*walletPublicKey]

	accessToken, err := twitterClient.Client.OAuthConsumer.AuthorizeToken(twitterClient.Client.OAuthTokens[tokenKey], verificationCode)
	if err != nil {
		http.Redirect(w, r, "https://triptychlabs.io", http.StatusTemporaryRedirect)
	}
	twitterClient.Client.CompleteAuth(accessToken)

	record := database.Record{
		PublicKey:        *walletPublicKey,
		AccessToken:      accessToken,
		VerificationCode: verificationCode,
		LastAuth:         time.Now().UTC().Unix(),
	}

	state.TwitterUsersPublicKey[accessToken.AdditionalData["screen_name"]] = *walletPublicKey

	record.WriteRecord()

	http.Redirect(w, r, "https://triptychlabs.io", http.StatusTemporaryRedirect)
}
