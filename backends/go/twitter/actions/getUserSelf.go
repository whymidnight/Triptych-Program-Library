package actions

import (
	"encoding/json"
	"log"

	"triptych.labs/twitter/v2/endpoints"
	"triptych.labs/twitter/v2/state"
	"triptych.labs/twitter/v2/structs"
)

func getUserSelfByPublicKey(publicKey string) {
}

func GetUserSelfByPublicKey(publicKey string) *structs.UserSelfResponse {
	var userSelfResponse structs.UserSelfResponse

	ident := state.TwitterClient[publicKey]

	data, err := ident.Client.BasicQuery(endpoints.USER_SELF)
	if err != nil {
		log.Println()
		return nil
	}

	err = json.Unmarshal(data, &userSelfResponse)
	if err != nil {
		log.Println(err)
		return nil
	}

	return &userSelfResponse
}
