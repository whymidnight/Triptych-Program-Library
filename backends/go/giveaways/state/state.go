package state

import (
	. "github.com/kkdai/twitter"
)

type Record struct {
	Client     *ServerClient
	LastAuth   int64
	OAuthToken string
}

type TwitterClientT map[string]*Record

var TwitterClient TwitterClientT

// TwitterAuthCodes[oauthToken] = publicKey
var TwitterAuthCodes map[string]*string

func Init() {
	TwitterClient = make(map[string]*Record)
	TwitterAuthCodes = make(map[string]*string)
}
