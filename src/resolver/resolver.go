package resolver

import (
	"github.com/diegofly91/apiturnos/src/service"
)

type Resolver struct {
	userService service.UserService
	authService service.AuthService
}
type ResolverDependencies struct {
	UserService service.UserService
	AuthService service.AuthService
}

// NewResolver returns a new instance of Resolver with the given UserService.
func GraphResolver() *Resolver {

	return &Resolver{
		// user service is initialized in the resolver
		userService: service.NewUserService(),
		authService: service.NewAuthService(),
	}
}
