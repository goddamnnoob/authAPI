package dto

type LoginResponse struct {
	AccessToken   string `json:"access_token"`
	ResponseToken string `json:"refresh_token,omitempty"`
}
