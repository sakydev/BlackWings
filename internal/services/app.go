package services

import (
	"BlackWings/internal/repositories"
	"BlackWings/internal/types"
	"context"
	"database/sql"

	"github.com/samber/do"
)

func InjectAppService(i *do.Injector) (*AppService, error) {
	return &AppService{
		appRepo: do.MustInvoke[*repositories.AppRepository](i),
	}, nil
}

type AppService struct {
	appRepo *repositories.AppRepository
}

func (s AppService) Create(ctx context.Context, database *sql.DB, options types.AppFlags) error {
	return nil
}
