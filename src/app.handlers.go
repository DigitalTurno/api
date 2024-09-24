package src

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"apiturnos/src/generated"
	"apiturnos/src/schema/directives"
	"apiturnos/src/schema/migration"
	"apiturnos/src/schema/resolver"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func PlaygroundHandler() *handler.Server {
	migration.MigrateTable()
	res := resolver.GraphResolver()
	c := generated.Config{Resolvers: res}
	c.Directives.Auth = directives.Auth
	c.Directives.HasRole = directives.HasRole

	srv := handler.New(generated.NewExecutableSchema(c))
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Puedes agregar lógica para permitir orígenes específicos
				return true // Permitir todos los orígenes para desarrollo
			},
		},
	})

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)
		var myErr *gqlerror.Error
		if errors.As(e, &myErr) {
			err.Message = myErr.Message
		}
		return err
	})
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		res := next(ctx)
		// Deserializar el json.RawMessage a un mapa
		var data map[string]interface{}
		if err := json.Unmarshal(res.Data, &data); err != nil {
			return res
		}
		for _, v := range data {
			// Asegurarse de que v es del tipo esperado
			if newData, ok := v.(map[string]interface{}); ok {
				if _, ok := newData["directives"]; ok {
					return res
				}
				var err error
				res.Data, err = json.Marshal(newData)
				if err != nil {
					return res
				}
			}
			break // Asumimos que solo hay un campo, por lo tanto, salimos del bucle
		}
		return res
	})
	// Soporte para consultas POST estándar
	return srv
}
