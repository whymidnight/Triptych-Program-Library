package helpers

/*
  This Package describes the helper functions to manage Authentication State

  Proposed scope:
    - (Re)Auth

    > Validate
      * Success
      * FAIL
        > Invalidate

  When Authentication is Invalidated,
    a `reauth` is required
*/

import (
	"time"

	"triptych.labs/twitter/v2/api"
	"triptych.labs/twitter/v2/constants"
	"triptych.labs/twitter/v2/database"
)

func Auth() (client *api.ServerClient, timestamp int64, requestUrl, oauthToken string) {
	client = api.NewServerClient(constants.ConsumerKey, constants.ConsumerSecret)

	timestamp = time.Now().UTC().Unix()

	requestUrl, oauthToken, _ = client.GetAuthURL(constants.CallbackURL)

	return
}

func ValidateAuth(publicKey string) (valid bool) {
	if record := database.FindAndReadRecord(publicKey); record != nil {
    /*
		if time.Now().Unix() >= record.LastAuth+60*60 {
			// TODO INVOKE TOKEN INVALIDATION WITH TWITTER
			// discard token
				if record.OAuthToken != "" {
					record.OAuthToken = ""
					record.WriteRecord()
				}

			valid = false
		} else {
			valid = true
		}
    */
    valid = true
	} else {
		valid = false
	}

	return
}

func InvalidateAuth() {}
