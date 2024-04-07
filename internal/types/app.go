package types

type App struct {
	ID       int64  `json:"ID"`
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

type AppFlags struct {
	Name string
}
