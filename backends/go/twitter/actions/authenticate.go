package actions

import (
	"triptych.labs/twitter/v2/helpers"
	"triptych.labs/twitter/v2/structs"
)

const PUBLIC_KEY = "publicKey"

type AuthResponse struct {
	Valid    bool                      `json:"valid"`
	UserData *structs.UserSelfResponse `json:"userData"`
}

func Authenticate(body interface{}) (response *AuthResponse, debug string) {
	var authResponse AuthResponse

	authRequest, ok := body.(map[string]interface{})
	if !ok {
		response = nil
		debug = "malformed method request body"
		return
	}

	publicKey := authRequest[PUBLIC_KEY].(string)
	authResponse.Valid = helpers.ValidateAuth(publicKey)
	if authResponse.Valid {
		authResponse.UserData = GetUserSelfByPublicKey(publicKey)
	}

	return &authResponse, ""
}
