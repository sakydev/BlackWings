package types

type Account struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	ClientID        string `json:"clientId"`
	ClientSecret    string `json:"clientSecret"`
	CredentialsJSON string `json:"CredentialsJson"`
	TokenJSON       string `json:"tokenJson"`
	App             App    `json:"app"`
}
