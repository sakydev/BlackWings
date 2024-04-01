package repositories

import (
	"context"
	"database/sql"

	"github.com/samber/do"
)

func InjectSettingRepository(i *do.Injector) (SettingRepository, error) {
	return SettingImpl{}, nil
}

type SettingImpl struct{}

type SettingRepository interface {
	List(ctx context.Context, database *sql.DB) (map[string]string, error)
}

func (impl SettingImpl) List(ctx context.Context, database *sql.DB) (map[string]string, error) {
	settings := make(map[string]string)

	rows, err := database.QueryContext(ctx, `
		SELECT name, value
		FROM settings
	`)
	if err != nil {
		return settings, err
	}
	defer rows.Close()

	for rows.Next() {
		var name, value string
		err = rows.Scan(&name, &value)
		if err != nil {
			return settings, err
		}

		settings[name] = value
	}

	return settings, nil
}
