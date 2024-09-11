package resolver

import (
	auth "apiturnos/src/modules/auth/service"
	profile "apiturnos/src/modules/profile/service"
	user "apiturnos/src/modules/user/service"
)

type Resolver struct {
	userService    user.UserService
	authService    auth.AuthService
	profileService profile.ProfileService
}
type ResolverDependencies struct {
	UserService    user.UserService
	AuthService    auth.AuthService
	ProfileService profile.ProfileService
}

// NewResolver returns a new instance of Resolver with the given UserService.
func GraphResolver() *Resolver {

	return &Resolver{
		// user service is initialized in the resolver
		userService:    user.NewUserService(),
		authService:    auth.NewAuthService(),
		profileService: profile.NewProfileService(),
	}
}
