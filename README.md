api/
├── gqlgen.yml
├── go.mod
├── main.go
└── src/
    └── modules/
        ├── user/
        |   ├── generated.go
        |   ├── models_gen.go
        |   ├── gql
        |   |   └── user.graphqls
        |   └── resolver/
        |       └── user_resolver.go
        ├── role/
        |   ├── generated.go
        |   ├── models_gen.go
        |   ├── gql
        |   |   └── role.graphqls
        |   └── resolver/
        |       └── role_resolver.go
        └── ...