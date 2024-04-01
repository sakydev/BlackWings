package types

import (
	"context"
	"database/sql"

	"github.com/samber/do"
)

type CommandConfiguration struct {
	Injector *do.Injector
	Context  context.Context
	Database *sql.DB
	Settings map[string]string
	Format   string
}
