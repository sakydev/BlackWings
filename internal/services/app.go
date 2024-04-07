package services

import (
	"black-wings/internal/repositories"
	"black-wings/internal/types"
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/samber/do"
)

func InjectAppService(i *do.Injector) (*AppService, error) {
	return &AppService{
		appRepo: do.MustInvoke[repositories.AppRepository](i),
	}, nil
}

type AppService struct {
	appRepo repositories.AppRepository
}

func (s AppService) GetByName(ctx context.Context, database *sql.DB, name string) (types.App, error) {
	return s.appRepo.GetByName(ctx, database, name)
}

func (s AppService) List(ctx context.Context, database *sql.DB) (map[string]types.App, error) {
	return s.appRepo.List(ctx, database)
}

func (s AppService) MapNamesToIDs(names []string, apps map[string]types.App) ([]int64, error) {
	var appIDs []int64

	for _, name := range names {
		matched := false
		for _, app := range apps {
			if strings.ToLower(app.Name) == strings.ToLower(name) {
				appIDs = append(appIDs, app.ID)
				matched = true

				break
			}

			if !matched {
				return appIDs, fmt.Errorf("app with name %s not found", name)
			}
		}
	}

	return appIDs, nil
}
