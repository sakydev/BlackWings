package repositories

import "github.com/samber/do"

func Wire(i *do.Injector) {
	do.Provide(i, InjectSettingRepository)
	do.Provide(i, InjectAccountRepository)
	do.Provide(i, InjectAppRepository)
}
