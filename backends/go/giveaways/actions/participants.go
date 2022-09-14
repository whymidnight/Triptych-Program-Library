package actions

import (
	"log"

	"triptych.labs/giveaways/v2/database"
)

type ParticipantsResponse struct {
	Giveaway     database.Giveaway `json:"giveaway"`
	Participants [][2]string       `json:"participants"`
}

func Participants(originalMessage string, body interface{}) (response *ParticipantsResponse, debug string) {
	var participantsResponse ParticipantsResponse

	participantsRequest, ok := body.(map[string]interface{})
	if !ok {
		response = nil
		debug = "malformed method request body"
		return
	}

	tweetId := participantsRequest[TWEET_ID].(string)
	log.Println(tweetId)
	giveaway := *database.FindRecord(tweetId)

	// sanitize giveaway profiles to not leak public keys
	var participants = make([][2]string, 0)
	for _, profile := range giveaway.Profiles {
		participants = append(participants, [2]string{profile[0], "image"})
	}
	giveaway.Profiles = make([][2]string, 0)

	participantsResponse.Giveaway = giveaway
	participantsResponse.Participants = participants

	return &participantsResponse, ""
}
