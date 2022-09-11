package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mrjones/oauth"
)

const DATA_LOCATION = "./storage/twitter"

type Record struct {
	PublicKey        string
	AccessToken      *oauth.AccessToken
	VerificationCode string
	LastAuth         int64
}

func (record *Record) WriteRecord() {
	// write $publicKey.json file of Record
	recordLocation := fmt.Sprintf("%s/%s.json", DATA_LOCATION, record.PublicKey)

	recordBytes, err := json.Marshal(record)
	if err != nil {
		log.Println(fmt.Sprintf("cannot marshal %s", recordLocation))
		return
	}

	ioutil.WriteFile(recordLocation, recordBytes, 0755)
}

func FindAndReadRecord(publicKey string) *Record {
	// glob data files for $publicKey.json
	var record Record

	recordLocation := fmt.Sprintf("%s/%s.json", DATA_LOCATION, publicKey)
	recordBytes, err := ioutil.ReadFile(recordLocation)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(recordBytes, &record)
	if err != nil {
		log.Println(fmt.Sprintf("cannot unmarshal %s", recordLocation))
		return nil
	}

	return &record
}
