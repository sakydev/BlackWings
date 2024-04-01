package internal

import (
	"BlackWings/internal/repositories"
	"BlackWings/internal/services"

	"github.com/samber/do"
)

func WireDependencies(i *do.Injector) {
	repositories.Wire(i)
	services.Wire(i)
}
