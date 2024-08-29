package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/diegofly91/apiturnos/src/middleware"
	"github.com/diegofly91/apiturnos/src/model"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {

	tokenData, err := middleware.CtxValue(ctx)
	if err != nil {
		return nil, err
	}
	if tokenData.Role != role {
		return nil, fmt.Errorf("Access Denied for this role")
	}
	// Puedes usar tokenData en los resolvers
	return next(ctx)
}
