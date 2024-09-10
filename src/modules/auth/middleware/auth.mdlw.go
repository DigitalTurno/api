package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"apiturnos/src/constants"
	"apiturnos/src/modules/auth/service"
	"apiturnos/src/schema/model"
)

type AuthMiddleware struct {
	authService service.AuthService
}

func NewAuthMiddleware() *AuthMiddleware {
	authService := service.NewAuthService()
	return &AuthMiddleware{
		authService: authService,
	}
}

func (am *AuthMiddleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		refreshToken := r.Header.Get("Refresh-Token")

		token := extractToken(authHeader)
		if token == "" {
			next.ServeHTTP(w, r)
			return
		}

		claims, err := am.authService.JwtValidate(context.Background(), token, constants.TypeToken("TOKEN"))
		if err != nil || !claims.Valid {
			if refreshToken != "" {
				token = extractToken(refreshToken)
				refreshClaims, err := am.authService.JwtValidate(context.Background(), token, constants.TypeToken("REFRESH"))

				if err != nil || !refreshClaims.Valid {
					// El refresh token tampoco es válido, rechazar la solicitud
					errorResponse(w, err)
					return
				}
				claims = refreshClaims

				// Si el refresh token es válido, generar un nuevo token JWT
				userPayload, _ := refreshClaims.Claims.(*model.UserPayload)
				newToken, err := am.authService.GenerateJWT(userPayload)
				if err != nil {
					errorResponse(w, err)
					return
				}

				// Devolver el nuevo token JWT en el encabezado
				w.Header().Set("auth-refresh", "Bearer "+newToken)
				errorResponse(w, fmt.Errorf("Update token JWT"))
				return

			} else {
				errorResponse(w, err)
				return
			}
			// Construir la estructura del error en formato JSON
		}

		customClaim, ok := claims.Claims.(*model.UserPayload)
		if !ok {
			errorResponse(w, fmt.Errorf("Invalid token data"))
			return
		}

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

func extractToken(authHeader string) string {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return ""
	}
	return strings.TrimPrefix(authHeader, "Bearer ")
}

func errorResponse(w http.ResponseWriter, err error) {
	errorResponse := map[string]interface{}{
		"errors": []map[string]interface{}{
			{
				"message": err.Error(),
			},
		},
	}
	// Establecer el encabezado de la respuesta como JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)

	// Enviar el error en formato JSON
	json.NewEncoder(w).Encode(errorResponse)
}
