package actions

import (
	"fmt"
	"log"
	"time"

	"triptych.labs/giveaways/v2/database"
	"triptych.labs/giveaways/v2/jobs"
	"triptych.labs/giveaways/v2/utils"
)

const TWEET_ID = "tweetId"
const EPOCH = "epoch"

type CreateResponse struct {
	Status bool `json:"status"`
}

func Create(originalMessage string, body interface{}) (response *CreateResponse, debug string) {
	var createResponse CreateResponse

	createRequest, ok := body.(map[string]interface{})
	if !ok {
		response = nil
		debug = "malformed method request body"
		return
	}

	publicKey := createRequest[PUBLIC_KEY].(string)
	tweetId := createRequest[TWEET_ID].(string)
	epochF := createRequest[EPOCH].(float64)
	epoch := int64(epochF)

	honest := utils.VerifyMessage(originalMessage, createRequest[originalMessage].(string), publicKey)
	if !honest {
		return nil, "bad identity"
	}

	if epoch <= time.Now().Unix() {
		epoch = time.Now().Unix()
	}

	if record := database.FindRecord(tweetId); record != nil {
		return nil, "Tweet already used"
	}

	log.Println("Waldo")
	jobs.AddJob(fmt.Sprintf("Tweet Id: %s", tweetId), jobs.Job{
		ExecutionTime: epoch,
		TweetId:       tweetId,
	})

	giveaway := database.Giveaway{
		PublicKey: publicKey,
		TweetId:   tweetId,
		StartTime: time.Now().Unix(),
		EndTime:   epoch,
		Winner:    [2]string{"", ""},
	}
	giveaway.WriteRecord(publicKey)

	createResponse.Status = true

	return &createResponse, ""
}
