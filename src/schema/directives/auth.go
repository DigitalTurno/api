package directives

import (
	"context"

	"apiturnos/src/modules/auth/middleware"

	"github.com/99designs/gqlgen/graphql"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, err := middleware.CtxValue(ctx)
	if err != nil {
		return nil, err
	}

	// Puedes usar tokenData en los resolvers
	return next(ctx)
}
