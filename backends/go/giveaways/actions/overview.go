package actions

import (
	"log"

	"triptych.labs/giveaways/v2/database"
	"triptych.labs/giveaways/v2/utils"
)

const PUBLIC_KEY = "publicKey"

type OverviewResponse struct {
	Valid     bool                `json:"valid"`
	Giveaways []database.Giveaway `json:"giveaways"`
}

func Overview(originalMessage string, body interface{}) (response *OverviewResponse, debug string) {
	var overviewResponse OverviewResponse

	overviewRequest, ok := body.(map[string]interface{})
	if !ok {
		response = nil
		debug = "malformed method request body"
		return
	}

	publicKey := overviewRequest[PUBLIC_KEY].(string)
	honest := utils.VerifyMessage(originalMessage, overviewRequest[originalMessage].(string), publicKey)
	if !honest {
		return nil, "bad identity"
	}

	log.Println("Waldo")

	overviewResponse.Valid = true
	overviewResponse.Giveaways = database.FindAndReadRecords(publicKey)
	log.Println("Walda")

	return &overviewResponse, ""
}
