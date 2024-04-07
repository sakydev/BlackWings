package internal

import (
	"black-wings/internal/repositories"
	"black-wings/internal/services"

	"github.com/samber/do"
)

func WireDependencies(i *do.Injector) {
	repositories.Wire(i)
	services.Wire(i)
}
