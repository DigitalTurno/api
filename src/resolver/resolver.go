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
func NewResolver(deps ResolverDependencies) *Resolver {
	return &Resolver{
		userService: deps.UserService,
	}
}
