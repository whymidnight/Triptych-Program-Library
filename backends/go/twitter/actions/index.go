package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"triptych.labs/twitter/v2/middleware"
	"triptych.labs/twitter/v2/structs"
)

func Handle(httpRequest *http.Request) (body interface{}, statusCode int) {
	var r structs.RequestT

	bodyBytes, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		body, statusCode = "invalid request", http.StatusBadRequest
		return
	}

	err = json.Unmarshal(bodyBytes, &r)
	if err != nil {
		body, statusCode = "malformed request", http.StatusBadRequest
		return
	}

	switch r.Method {
	case "authenticate":
		{
			authResponse, debug := Authenticate(r.Body)
			if authResponse == nil {
				return debug, 400
			}

			return authResponse, http.StatusOK
		}
	}

	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	middleware.SetupResponse(&w)

	var response = new(structs.ResponseT)

	switch (*r).Method {
	case "OPTIONS":
		return
	case "POST":
		{
			body, statusCode := Handle(r)

			response.WriteNewResponse(&w, body, statusCode)
			return
		}
	}
}
