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
	Create(ctx context.Context, database *sql.DB, app types.App, accountDetails types.Account) (int64, error)
}

func (impl AccountImpl) Create(ctx context.Context, database *sql.DB, app types.App, accountDetails types.Account) (int64, error) {
	accountID := int64(0)
	rows, err := database.ExecContext(ctx, `
		INSERT INTO accounts
		(client_id, client_secret, raw, app_id)
		VALUES ($1, $2, $3, $4)
	`, accountDetails.ClientID, accountDetails.ClientSecret, accountDetails.Raw, app.ID)
	if err != nil {
		return accountID, err
	}

	accountID, err = rows.LastInsertId()
	if err != nil {
		return accountID, err
	}

	return accountID, nil
}
