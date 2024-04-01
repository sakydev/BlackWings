package types

type Account struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	Raw          string `json:"raw"`
	App          App    `json:"app"`
}
