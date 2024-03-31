package internal

import (
	"BlackWings/internal/services"

	"github.com/samber/do"
)

func WireDependencies(i *do.Injector) {
	services.Wire(i)
}
