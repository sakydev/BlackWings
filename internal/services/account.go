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
