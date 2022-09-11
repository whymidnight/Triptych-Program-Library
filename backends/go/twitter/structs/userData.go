package structs

type UserSelfData struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type UserSelfResponse struct {
	Data UserSelfData `json:"data"`
}
