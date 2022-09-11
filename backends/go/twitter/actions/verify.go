package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"triptych.labs/twitter/v2/helpers"
	"triptych.labs/twitter/v2/middleware"
	"triptych.labs/twitter/v2/state"
	"triptych.labs/twitter/v2/structs"

	"triptych.labs/utils"
)

type VerifyResponse struct {
	RedirectTo string `json:"redirectTo"`
}

func Verify(w http.ResponseWriter, r *http.Request) {
	var response = new(structs.ResponseT)

	middleware.SetupResponse(&w)

	switch (*r).Method {
	case "OPTIONS":
		return
	case "POST":
		{
			walletPublicKey := r.Header.Get("PublicKey")
			originalMessage := r.Header.Get("OriginalMessage")

			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				response.WriteNewResponse(&w, "bad body in request", http.StatusBadRequest)
				return
			}

			var body map[string]string
			err = json.Unmarshal(bodyBytes, &body)
			if err != nil {
				response.WriteNewResponse(&w, "malformed body in request", http.StatusBadRequest)
				return
			}

			honest := utils.VerifyMessage(originalMessage, body[originalMessage], walletPublicKey)
			if !honest {
				response.WriteNewResponse(&w, "bad identity", http.StatusUnauthorized)
				return
			}

			client, _, requestUrl, oauthToken := helpers.Auth()

			_, ok := state.TwitterClient[walletPublicKey]
			if !ok {
				state.TwitterClient[walletPublicKey] = &state.Record{}
			}

			state.TwitterAuthCodes[oauthToken] = &walletPublicKey
			state.TwitterClient[walletPublicKey].Client = client
			state.TwitterClient[walletPublicKey].RequestUrl = requestUrl

			response.WriteNewResponse(&w, VerifyResponse{requestUrl}, 200)
			return
		}
	}
}


