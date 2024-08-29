package directives

import (
	"context"
	"fmt"

	"apiturnos/src/modules/auth/middleware"
	"apiturnos/src/schema/model"

	"github.com/99designs/gqlgen/graphql"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, roles []model.Role) (interface{}, error) {

	tokenData, err := middleware.CtxValue(ctx)
	if err != nil {
		return nil, err
	}
	roleAllowed := false
	for _, role := range roles {
		if tokenData.Role == role {
			roleAllowed = true
			break
		}
	}
	if !roleAllowed {
		return nil, fmt.Errorf("Access Denied for you role")
	}
	// Puedes usar tokenData en los resolvers
	return next(ctx)
}
