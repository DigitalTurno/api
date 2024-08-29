package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/diegofly91/apiturnos/src/middleware"
	"github.com/diegofly91/apiturnos/src/model"
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
