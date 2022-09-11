package constants

import "os"

const (
	CallbackURL string = "https://triptychlabs.io:43594/twitter/maketoken"
)

var ConsumerKey string
var ConsumerSecret string

func Init() {
	ConsumerKey = os.Getenv("ConsumerKey")
	ConsumerSecret = os.Getenv("ConsumerSecret")
}
