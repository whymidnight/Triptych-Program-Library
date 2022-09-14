package database

import (
	"encoding/json"
	"log"

	"github.com/xujiajun/nutsdb"
)

const DATA_LOCATION = "./storage/giveaways"

const GIVEAWAYS_BUCKET = "bucket_giveawaysRecords"
const PUBLICKEY_TO_TWEET_IDS_BUCKET = "bucket_publicKeyToTweetId"

var DB *nutsdb.DB = nil

/*
  Bucket: "Giveaways"

    K: PublicKey
    V: {
      [tweetId]: {
        winner: string,
        startTime: int64
        endTime: int64
      }
    }
*/

type Giveaway struct {
	PublicKey    string
	TweetId      string
	Winner       [2]string
	StartTime    int64
	EndTime      int64
	Participants int64
	Hash         string
	Profiles     [][2]string
}

// map[tweetId]Giveaway
type TweetGiveaways map[string]Giveaway

func Init() {
	db, err := nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(DATA_LOCATION),
	)
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}

func (g *Giveaway) WriteRecord(publicKey string) {
	gJson, err := json.Marshal(g)
	if err != nil {
		return
	}

	tweetIdBytes := []byte(g.TweetId)

	if err := DB.Update(
		func(tx *nutsdb.Tx) error {

			size, err := tx.LSize(PUBLICKEY_TO_TWEET_IDS_BUCKET, []byte(publicKey))
			if err != nil {
				size = 0
			}
			log.Println(publicKey, size)
			if size != 0 {
				exists := false

				if items, err := tx.LRange(PUBLICKEY_TO_TWEET_IDS_BUCKET, []byte(publicKey), 0, -1); err != nil {
					log.Println("bad lrange", publicKey)
					return nil
				} else {
					for _, item := range items {
						if e, err := tx.Get(GIVEAWAYS_BUCKET, item); err != nil {
							log.Println("bad get", string(item))
							return nil
						} else {
							var giveaway Giveaway
							_ = json.Unmarshal(e.Value, &giveaway)
							if g.TweetId == giveaway.TweetId {
								exists = true
							}
						}
					}
				}
				if exists {
					log.Println("tweet re-entrance", g.TweetId)
					return nil
				}
			}

			if err := tx.RPush(PUBLICKEY_TO_TWEET_IDS_BUCKET, []byte(publicKey), tweetIdBytes); err != nil {
				log.Println("bad push")
				return nil
			}

			if err := tx.Put(GIVEAWAYS_BUCKET, tweetIdBytes, gJson, 0); err != nil {
				log.Println("bad put")
				return nil
			}

			return nil
		}); err != nil {
		log.Println("Bad Write Record")
	}

}

func FindAndReadRecords(publicKey string) []Giveaway {
	var giveaways = make([]Giveaway, 0)

	if err := DB.View(
		func(tx *nutsdb.Tx) error {

			size, err := tx.LSize(PUBLICKEY_TO_TWEET_IDS_BUCKET, []byte(publicKey))
			if err != nil {
				return nil
			}
			if size == 0 {
				return nil
			}

			if items, err := tx.LRange(PUBLICKEY_TO_TWEET_IDS_BUCKET, []byte(publicKey), 0, -1); err != nil {
				log.Println("bad lrange", publicKey)
				return nil
			} else {
				for _, item := range items {
					if e, err := tx.Get(GIVEAWAYS_BUCKET, item); err != nil {
						log.Println("bad get", string(item))
						return nil
					} else {
						var giveaway Giveaway
						_ = json.Unmarshal(e.Value, &giveaway)
						giveaways = append(giveaways, giveaway)
					}
				}
			}

			return nil
		}); err != nil {
		log.Println("bad reads")
	}

	return giveaways
}

func FindRecord(tweetId string) *Giveaway {
	var giveaway *Giveaway = nil

	if err := DB.View(
		func(tx *nutsdb.Tx) error {
			if e, err := tx.Get(GIVEAWAYS_BUCKET, []byte(tweetId)); err != nil {
				return nil
			} else {
				_ = json.Unmarshal(e.Value, &giveaway)
			}

			return nil
		}); err != nil {
		log.Println("bad FindRecord")
	}

	return giveaway
}

func (g *Giveaway) UpdateRecord() {
	updateJs, _ := json.Marshal(g)

	if err := DB.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.Put(GIVEAWAYS_BUCKET, []byte(g.TweetId), updateJs, 0); err != nil {
				return nil
			}
			return nil
		}); err != nil {
		log.Println("bad FindRecord")
	}

}
