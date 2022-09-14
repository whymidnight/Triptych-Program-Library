package actions

import (
	"encoding/json"
	"fmt"
	"log"

	"triptych.labs/twitter/v2/api"
	"triptych.labs/twitter/v2/helpers"
	"triptych.labs/twitter/v2/state"
)

type TweetLikesResponse struct {
	NumberOfLikes    int `json:"numberOfLikes"`
	Profiles [][2]string `json:"numberOfProfiles"`
}

func GetTweet(tweetId string) *TweetLikesResponse {
	var response TweetLikesResponse

	type LikedUser struct {
		Id       string `json:"id"`
		Name     string `json:"name"`
		Username string `json:"username"`
	}

	type ResultCount struct {
		NextToken *string `json:"next_token"`
	}

	type TweetLikesResponse struct {
		Data []LikedUser `json:"data,omitempty"`
		Meta ResultCount `json:"meta"`
	}

	usernames := make([]string, 0)

	var nextToken = ""
	for {
		tweetLikesUrl := fmt.Sprintf("https://api.twitter.com/2/tweets/%s/liking_users", tweetId)
		if nextToken != "" {
			tweetLikesUrl = fmt.Sprintf("%s?pagination_token=%s", tweetLikesUrl, nextToken)
		}

		var tweetLikesResponse TweetLikesResponse
		likes, err := api.AppRequest(state.AppBearerToken, tweetLikesUrl)
		if err != nil {
			log.Println(err)
			return nil
		}
		err = json.Unmarshal(likes, &tweetLikesResponse)
		if err != nil {
			log.Println(err)
			return nil
		}
		if tweetLikesResponse.Meta.NextToken == nil {
			break
		} else {
			// numberOfLikes += len(tweetLikesResponse.Data)
			for _, user := range tweetLikesResponse.Data {
				usernames = append(usernames, user.Username)
			}
			nextToken = *tweetLikesResponse.Meta.NextToken
		}
	}

	profiles := helpers.FindPublicKeysfromUsernames(usernames)
	response.NumberOfLikes = len(usernames)
  response.Profiles = profiles
  

	return &response
}
