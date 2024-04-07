package repositories

import (
	"black-wings/internal/types"
	"context"
	"database/sql"

	"github.com/samber/do"
)

func InjectAppRepository(i *do.Injector) (AppRepository, error) {
	return AppImpl{}, nil
}

type AppImpl struct{}

type AppRepository interface {
	GetByName(ctx context.Context, database *sql.DB, name string) (types.App, error)
	List(ctx context.Context, database *sql.DB) (map[string]types.App, error)
}

func (impl AppImpl) GetByName(ctx context.Context, database *sql.DB, name string) (types.App, error) {
	app := types.App{}

	row := database.QueryRowContext(ctx, `
		SELECT id, name, provider
		FROM apps
		WHERE name = $1
	`, name)

	err := row.Scan(&app.ID, &app.Name, &app.Provider)
	if err != nil {
		return app, err
	}

	return app, nil
}

func (impl AppImpl) List(ctx context.Context, database *sql.DB) (map[string]types.App, error) {
	apps := make(map[string]types.App)

	rows, err := database.QueryContext(ctx, `
		SELECT id, name, provider
		FROM apps
	`)
	if err != nil {
		return apps, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var name, provider string
		err = rows.Scan(&id, &name, &provider)
		if err != nil {
			return apps, err
		}

		apps[name] = types.App{ID: id, Name: name, Provider: provider}
	}

	return apps, nil
}
