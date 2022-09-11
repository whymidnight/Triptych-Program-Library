package api

import (
	"net/http"

	"github.com/mrjones/oauth"
)

type ServerClient struct {
	HttpConn      *http.Client
	OAuthConsumer *oauth.Consumer
	OAuthTokens   map[string]*oauth.RequestToken
}
