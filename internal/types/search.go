package types

type SearchFlags struct {
	Query     string
	Apps      []string
	Include   string
	Exclude   string
	FileTypes []string
	Before    string
	After     string
	Sort      string
	Order     string
	Limit     int64
	Offset    int64
}

type SearchResult struct {
	Account     string `json:"account"`
	Service     string `json:"service"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Link        string `json:"link"`
	Type        string `json:"type"`
}
