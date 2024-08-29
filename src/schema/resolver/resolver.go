package resolver

import (
	auth "apiturnos/src/modules/auth/service"
	user "apiturnos/src/modules/user/service"
)

type Resolver struct {
	userService user.UserService
	authService auth.AuthService
}
type ResolverDependencies struct {
	UserService user.UserService
	AuthService auth.AuthService
}

// NewResolver returns a new instance of Resolver with the given UserService.
func GraphResolver() *Resolver {

	return &Resolver{
		// user service is initialized in the resolver
		userService: user.NewUserService(),
		authService: auth.NewAuthService(),
	}
}
