package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/diegofly91/apiturnos/src/middleware"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, err := middleware.CtxValue(ctx)
	if err != nil {
		return nil, err
	}

	// Puedes usar tokenData en los resolvers
	return next(ctx)
}
