package module

import (
	"github.com/diegofly91/apiturnos/src/resolver"
)

func SetupAppModule() *resolver.Resolver {
	// Initialize dependencies for the resolver.
	deps := resolver.ResolverDependencies{
		UserService: InitializeUserModule(),
	}
	res := resolver.NewResolver(deps)
	return res
}
