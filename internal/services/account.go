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

func (s AccountService) Connect(ctx context.Context, database *sql.DB, options types.AppFlags) (string, error) {
	return "", nil
}
