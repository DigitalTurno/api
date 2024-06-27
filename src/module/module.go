package module

import (
	db "github.com/diegofly91/apiturnos/src/config"
	"github.com/diegofly91/apiturnos/src/resolver"
)

func SetupModule() *resolver.Resolver {
	db := db.Database
	userService := InitializeUserModule(db)
	deps := resolver.ResolverDependencies{
		UserService: userService,
	}
	res := resolver.NewResolver(deps)
	return res
}
