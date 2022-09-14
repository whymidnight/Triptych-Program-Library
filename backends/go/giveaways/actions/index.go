package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"triptych.labs/giveaways/v2/middleware"
	"triptych.labs/giveaways/v2/structs"
)

func Handle(httpRequest *http.Request) (body interface{}, statusCode int) {
	var r structs.RequestT

	message := httpRequest.Header.Get("OriginalMessage")

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

	var response interface{}
	var debug = ""

	switch r.Method {
	case "overview":
		{

			response, debug = Overview(message, r.Body)
		}
	case "create":
		{

			response, debug = Create(message, r.Body)
		}
	case "rollInvoke":
		{

			response, debug = RollInvoke(message, r.Body)
		}
	case "rollStart":
		{

			response, debug = RollStart(message, r.Body)
		}
	case "rollConfirm":
		{

			response, debug = RollConfirm(message, r.Body)
		}
	case "participants":
		{

			response, debug = Participants(message, r.Body)
		}
	default:
		{
			response = "unsupported"
		}
	}
	if response == nil {
		return debug, 400
	}

	return response, 200
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
