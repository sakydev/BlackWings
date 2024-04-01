package repositories

import (
	"BlackWings/internal/types"
	"context"
	"database/sql"

	"github.com/samber/do"
)

func InjectAccountRepository(i *do.Injector) (AccountRepository, error) {
	return AccountImpl{}, nil
}

type AccountImpl struct{}

type AccountRepository interface {
	Create(ctx context.Context, database *sql.DB, app types.App) (int64, error)
}

func (impl AccountImpl) Create(ctx context.Context, database *sql.DB, app types.App) (int64, error) {
	appID := int64(0)
	rows, err := database.ExecContext(ctx, `
		INSERT INTO apps
		(name, provider)
		VALUES ($1, $2)
	`, app.Name, app.Provider)
	if err != nil {
		return appID, err
	}

	appID, err = rows.LastInsertId()
	if err != nil {
		return appID, err
	}

	return appID, nil
}
