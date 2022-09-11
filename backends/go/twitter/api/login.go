package api

import (
	"log"

	"github.com/mrjones/oauth"
)

const (
	OAUTH_REQUES_TOKEN string = "https://api.twitter.com/oauth/request_token"
	OAUTH_AUTH_TOKEN   string = "https://api.twitter.com/oauth/authorize"
	OAUTH_ACCESS_TOKEN string = "https://api.twitter.com/oauth/access_token"
)

func NewServerClient(consumerKey, consumerSecret string) *ServerClient {
	newServer := new(ServerClient)

	newServer.OAuthConsumer = oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   OAUTH_REQUES_TOKEN,
			AuthorizeTokenUrl: OAUTH_AUTH_TOKEN,
			AccessTokenUrl:    OAUTH_ACCESS_TOKEN,
		},
	)
	newServer.OAuthConsumer.Debug(false)
	newServer.OAuthTokens = make(map[string]*oauth.RequestToken)

	return newServer
}

func (s *ServerClient) GetAuthURL(tokenUrl string) (string, string, error) {
	token, requestUrl, err := s.OAuthConsumer.GetRequestTokenAndUrl(tokenUrl)

	if err != nil {
		log.Println(err)
		return "", "", nil
	}
	s.OAuthTokens[token.Token] = token
	return requestUrl, token.Token, nil
}

func (s *ServerClient) CompleteAuth(accessToken *oauth.AccessToken) error {
	var err error
	s.HttpConn, err = s.OAuthConsumer.MakeHttpClient(accessToken)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
