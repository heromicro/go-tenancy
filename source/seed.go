package source

import "github.com/snowlyg/go-tenancy/model"

// RunSeed
func RunSeed() error {
	return model.Seed(
		Admin,
		Api,
		AuthorityMenu,
		Authority,
		AuthoritiesMenus,
		Casbin,
		DataAuthorities,
		BaseMenu,
		Region,
		Config,
		SysConfigCategory,
		SysConfigValue,
		FinancialRecord,
	)
}
