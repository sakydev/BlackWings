package internal

import (
	"black-wings/internal/repositories"
	"black-wings/internal/services"
	"black-wings/internal/services/integrations"

	"github.com/samber/do"
)

func WireDependencies(i *do.Injector) {
	repositories.Wire(i)
	integrations.Wire(i)
	services.Wire(i)
}
