package repositories

import (
	"BlackWings/internal/types"
	"BlackWings/internal/utils"
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
	List(ctx context.Context, database *sql.DB) ([]types.Account, error)
	ListByApps(ctx context.Context, database *sql.DB, appIDs []int64) ([]types.Account, error)
}

func (impl AccountImpl) Create(ctx context.Context, database *sql.DB, app types.App, accountDetails types.Account) (int64, error) {
	accountID := int64(0)
	rows, err := database.ExecContext(ctx, `
		INSERT INTO accounts
		(name, client_id, client_secret, raw, app_id)
		VALUES ($1, $2, $3, $4, $5)
	`, accountDetails.Name, accountDetails.ClientID, accountDetails.ClientSecret, accountDetails.Raw, app.ID)
	if err != nil {
		return accountID, err
	}

	accountID, err = rows.LastInsertId()
	if err != nil {
		return accountID, err
	}

	return accountID, nil
}

func (impl AccountImpl) List(ctx context.Context, database *sql.DB) ([]types.Account, error) {
	var accounts []types.Account

	rows, err := database.QueryContext(ctx, `
		SELECT ac.name, ac.client_id, ac.client_secret, ac.raw, ap.id, ap.name, ap.provider
		FROM accounts ac
		INNER JOIN apps ap ON ac.app_id = ap.id
	`)
	if err != nil {
		return accounts, err
	}
	defer rows.Close()

	accounts, err = processAccountRows(rows)

	return accounts, nil
}

func (impl AccountImpl) ListByApps(ctx context.Context, database *sql.DB, appIDs []int64) ([]types.Account, error) {
	var accounts []types.Account

	rows, err := database.QueryContext(ctx, `
		SELECT ac.name, ac.client_id, ac.client_secret, ac.raw, ap.id, ap.name, ap.provider
		FROM accounts ac
		INNER JOIN apps ap ON ac.app_id = ap.id
		WHERE ap.id IN ($1)
	`, utils.ImplodeInt64(appIDs, ","))
	if err != nil {
		return accounts, err
	}
	defer rows.Close()

	accounts, err = processAccountRows(rows)

	return accounts, err
}

func processAccountRows(rows *sql.Rows) ([]types.Account, error) {
	var accounts []types.Account

	for rows.Next() {
		var currentAccount types.Account
		var appID int
		var accountName, clientID, clientSecret, raw, appName, provider string

		err := rows.Scan(&accountName, &clientID, &clientSecret, &raw, &appID, &appName, &provider)
		if err != nil {
			return accounts, err
		}

		currentAccount = types.Account{
			Name:         accountName,
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Raw:          raw,
			App: types.App{
				ID:       int64(appID),
				Name:     appName,
				Provider: provider,
			},
		}

		accounts = append(accounts, currentAccount)
	}

	return accounts, nil
}
