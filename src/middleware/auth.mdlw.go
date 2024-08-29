package middleware

import (
	"context"
	"net/http"

	"github.com/diegofly91/apiturnos/src/constants"
	"github.com/diegofly91/apiturnos/src/model"
	"github.com/diegofly91/apiturnos/src/service"
)

type authString string

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			next.ServeHTTP(w, r)
			return
		}
		bearer := "Bearer "
		token = token[len(bearer):]

		authService := service.NewAuthService()
		validate, err := authService.JwtValidate(context.Background(), token)
		if err != nil || !validate.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}
		customClaim, _ := validate.Claims.(*model.UserPayload)

		ctx := context.WithValue(r.Context(), constants.TokenDataKey, customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) *model.UserPayload {
	raw, _ := ctx.Value(constants.TokenDataKey).(*model.UserPayload)
	return raw
}
