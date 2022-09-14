package structs

type UserSelfData struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type UserSelfResponse struct {
	Data UserSelfData `json:"data"`
}
