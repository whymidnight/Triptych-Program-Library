package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (c *ServerClient) BasicQuery(queryString string) ([]byte, error) {
	if c.HttpConn == nil {
		return nil, errors.New("No Client OAuth")
	}

	response, err := c.HttpConn.Get(queryString)
	if err != nil {
		log.Println(err)
		return []byte{}, errors.New("bad response")
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	return bits, err
}

func AppRequest(appBearerToken, queryString string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", queryString, nil)
	if err != nil {
		return []byte{}, errors.New("bad request")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", appBearerToken))

	res, err := client.Do(req)
	if err != nil {
		return []byte{}, errors.New("bad response")
	} else {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, errors.New("bad response body")
	}

	return body, nil
}
