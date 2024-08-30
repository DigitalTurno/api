package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"apiturnos/src/constants"
	"apiturnos/src/modules/auth/service"
	"apiturnos/src/schema/model"
)

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
			// Construir la estructura del error en formato JSON
			errorResponse := map[string]interface{}{
				"errors": []map[string]interface{}{
					{
						"message": "Validate Token Error",
					},
				},
			}
			// Establecer el encabezado de la respuesta como JSON
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)

			// Enviar el error en formato JSON
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		customClaim, _ := validate.Claims.(*model.UserPayload)

		ctx := context.WithValue(r.Context(), constants.TokenDataKey, customClaim)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func CtxValue(ctx context.Context) (*model.UserPayload, error) {
	raw, ok := ctx.Value(constants.TokenDataKey).(*model.UserPayload)
	if !ok {
		return nil, fmt.Errorf("Invalid token data")
	}
	return raw, nil
}
