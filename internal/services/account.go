package services

import (
	"BlackWings/internal/repositories"
	"BlackWings/internal/types"
	"context"
	"database/sql"

	"github.com/samber/do"
)

func InjectAccountService(i *do.Injector) (*AccountService, error) {
	return &AccountService{
		accountRepo: do.MustInvoke[repositories.AccountRepository](i),
	}, nil
}

type AccountService struct {
	accountRepo repositories.AccountRepository
}

func (s AccountService) Connect(ctx context.Context, database *sql.DB, app types.App, accountDetails types.Account) (int64, error) {
	return s.accountRepo.Create(ctx, database, app, accountDetails)
}

func (s AccountService) List(ctx context.Context, database *sql.DB, appIDs []int64) ([]types.Account, error) {
	if len(appIDs) == 0 {
		return s.accountRepo.List(ctx, database)
	}

	return s.accountRepo.ListByApps(ctx, database, appIDs)
}
