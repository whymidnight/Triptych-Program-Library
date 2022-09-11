package structs

type RequestT struct {
	Method string      `json:"method"`
	Body   interface{} `json:"body"`
}
