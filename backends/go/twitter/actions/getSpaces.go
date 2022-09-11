package actions

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"triptych.labs/twitter/v2/api"
	"triptych.labs/twitter/v2/state"
)

type SpaceResponse struct {
	Id               string `json:"id"`
	ParticipantCount int    `json:"participantCount"`
	State            string `json:"state"`
	Title            string `json:"title"`
}

func GetSpace(userId string) *SpaceResponse {
	type SpacesResponse struct {
		Data []SpaceResponse `json:"data"`
	}
	var response SpaceResponse
	fields := []string{"title", "participant_count"}

	activeSpace := fmt.Sprintf("%s?%s=%s&space.fields=%s", "https://api.twitter.com/2/spaces/by/creator_ids", "user_ids", userId, strings.Join(fields, ","))

	var spacesResponse SpacesResponse
	spaces, err := api.AppRequest(state.AppBearerToken, activeSpace)
	if err != nil {
		log.Println(err)
		return nil
	}
	err = json.Unmarshal(spaces, &spacesResponse)
	if err != nil {
		log.Println(err)
		return nil
	}

	response = spacesResponse.Data[0]

	return &response
}
