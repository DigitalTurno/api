# Nombre del binario
BINARY_NAME=apiturnos

# Directorio donde se encuentran los archivos .graphql
GRAPHQL_DIR=schema

# Ruta al archivo gqlgen.yml
GQLGEN_CONFIG=gqlgen.yml

# Comandos
all: build

# Instala las dependencias
deps:
	go mod tidy

# Genera el código GraphQL a partir de los esquemas
generate:
	go run github.com/99designs/gqlgen generate

# Construye el binario
build: generate
	go build -o $(BINARY_NAME) ./server.go

# Ejecuta las pruebas
test:
	go test ./...

# Ejecuta el binario
run:
	go run server.go

# Limpia los archivos generados
clean:
	rm -f $(BINARY_NAME)

# Vuelve a generar el código y compilar
rebuild: clean all

.PHONY: all deps generate build test run clean rebuild
