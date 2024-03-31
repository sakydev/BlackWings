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
