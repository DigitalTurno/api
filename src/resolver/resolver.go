package resolver

import (
	"github.com/diegofly91/apiturnos/src/service"
)

type Resolver struct {
	userService service.UserService
}
type ResolverDependencies struct {
	UserService service.UserService
}

// NewResolver returns a new instance of Resolver with the given UserService.
func GraphResolver() *Resolver {

	return &Resolver{
		// user service is initialized in the resolver
		userService: service.NewUserService(),
	}
}
