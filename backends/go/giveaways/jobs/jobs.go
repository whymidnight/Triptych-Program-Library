package jobs

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/didil/goblero/pkg/blero"
	"triptych.labs/giveaways/v2/database"
	twitterActions "triptych.labs/twitter/v2/actions"
)

const DATA_LOCATION = "./storage/jobs"

var BL *blero.Blero = nil

type Job struct {
	ExecutionTime int64  `json:"executionTime"`
	TweetId       string `json:"tweetId"`
}

func Init() {
	bl := blero.New(DATA_LOCATION)
	err := bl.Start()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup

	// not really sure how else to rework the blero lib
	// 10k processors should be enough
	// hypothetically 10k giveaway concurrent tasks at once
	for i := 1; i <= 10000; i++ {
		pI := i
		go func() {
			wg.Add(1)
			bl.RegisterProcessorFunc(func(j *blero.Job) error {
				var job Job
				err := json.Unmarshal(j.Data, &job)
				if err != nil {
					return err
				}

				time.Sleep(time.Duration(job.ExecutionTime-time.Now().Unix()) * time.Second)

				log.Printf("[Processor %v] Processing job: %v - data: %v\n", pI, j.Name, string(j.Data))

				kpis := twitterActions.GetTweet(job.TweetId)
				giveaway := database.FindRecord(job.TweetId)

				giveaway.Participants = int64(kpis.NumberOfProfiles)
				giveaway.UpdateRecord()

				log.Printf("[Processor %v] Done Processing job: %v\n", pI, j.Name)

				return nil
			})
			wg.Done()
		}()
	}

	wg.Wait()

	BL = bl

	log.Println("blero init complete")
}

func Stop() {
	BL.Stop()
}

func AddJob(jobName string, jobData Job) error {
	jobDataJson, err := json.Marshal(jobData)
	if err != nil {
		return err
	}

	_, err = BL.EnqueueJob(jobName, jobDataJson)
	if err != nil {
		return err
	}

	return nil
}
