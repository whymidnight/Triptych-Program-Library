package state

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"triptych.labs/twitter/v2/api"
	"triptych.labs/twitter/v2/constants"
	"triptych.labs/twitter/v2/database"
)

type Record struct {
	Client     *api.ServerClient
	RequestUrl string
	Username   string
}

type TwitterClientT map[string]*Record

var AppBearerToken string
var TwitterClient TwitterClientT

// TwitterAuthCodes[oauthToken] = publicKey
var TwitterAuthCodes map[string]*string

var TwitterUsersPublicKey map[string]string

func Init() {
	TwitterClient = make(map[string]*Record)
	TwitterAuthCodes = make(map[string]*string)
	TwitterUsersPublicKey = make(map[string]string)

	// hydrate memory from database
	Hydrate()
	Setup()
}

func Hydrate() {
	log.Println("Hydrating...")
	files, err := ioutil.ReadDir(database.DATA_LOCATION)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		var record database.Record

		if !strings.Contains(file.Name(), ".json") {
			continue
		}

		path := fmt.Sprintf("%s/%s", database.DATA_LOCATION, file.Name())
		log.Println(path)
		recordBytes, _ := ioutil.ReadFile(path)

		_ = json.Unmarshal(recordBytes, &record)
		client := api.NewServerClient(constants.ConsumerKey, constants.ConsumerSecret)
		client.CompleteAuth(record.AccessToken)
		TwitterClient[record.PublicKey] = &Record{}
		TwitterClient[record.PublicKey].Client = client

		fmt.Println(record.AccessToken)

		TwitterClient[record.PublicKey].Username = record.AccessToken.AdditionalData["screen_name"]

		TwitterUsersPublicKey[TwitterClient[record.PublicKey].Username] = record.PublicKey

		log.Println(TwitterClient[record.PublicKey].Username)
	}
}

// fetch oauth2 app-only bearer token
func Setup() {
	type AppBearerTokenResponse struct {
		TokenType   string `json:"token_type"`
		AccessToken string `json:"access_token"`
	}
	bearerTokenCredential := fmt.Sprintf("%s:%s", constants.ConsumerKey, constants.ConsumerSecret)
	bearerTokenCredentialB64 := base64.RawStdEncoding.EncodeToString([]byte(bearerTokenCredential))

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.twitter.com/oauth2/token", bytes.NewBufferString("grant_type=client_credentials"))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", bearerTokenCredentialB64))

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var appBearerTokenResponse AppBearerTokenResponse
	err = json.Unmarshal(body, &appBearerTokenResponse)
	if err != nil {
		panic(err)
	}
	log.Println("App-Only bearer token", appBearerTokenResponse.AccessToken)
	AppBearerToken = appBearerTokenResponse.AccessToken
}
