package services

import (
	"black-wings/internal/services/apps"

	"github.com/samber/do"
)

func Wire(i *do.Injector) {
	do.Provide(i, apps.InjectGmailService)
	do.Provide(i, InjectAppService)
	do.Provide(i, InjectAccountService)
	do.Provide(i, InjectSearchService)
}
