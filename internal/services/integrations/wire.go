package integrations

import (
	"github.com/samber/do"
)

func Wire(i *do.Injector) {
	do.Provide(i, InjectGoogleService)
}
