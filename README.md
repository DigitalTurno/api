# API Project

## Table of Contents

- [Introduction](#introduction)
- [Architecture Overview](#architecture-overview)
- [Directory Structure](#directory-structure)
- [Setup and Installation](#setup-and-installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Welcome to the **API Project**! This project is a Go-based API leveraging GraphQL with the help of **gqlgen**. It's structured to promote scalability, maintainability, and clarity by organizing code into well-defined directories and modules. Whether you're contributing to the project or setting it up locally, this README will guide you through the architecture and setup.

## Architecture Overview

The project follows a modular architecture, separating concerns into distinct layers and components. This separation ensures that each part of the application is manageable, testable, and maintainable. The main components include:

- **Configuration**: Manages application configurations and database connections.
- **GraphQL Schema and Resolvers**: Defines the GraphQL schema and associated resolver logic.
- **Middlewares**: Handles cross-cutting concerns like logging, authentication, etc.
- **Models**: Defines data structures and interacts with the database.
- **Repositories**: Abstracts data access logic.
- **Services**: Contains business logic.
- **Utilities**: Provides helper functions and utilities used across the application.

## Directory Structure

Here's an overview of the project's directory structure:

```
api/
├── gqlgen.yml
├── go.mod
├── .env
├── server.go
└── src/
    ├── app.handlers.go
    ├── config/
    |   └── database.go
    ├── generate/
    |   └── generate.go
    ├── modules/
    |   └── {name}/
    |       ├── repository/
    |       |   └── **.repo.go
    |       ├── service/
    |       |   └── **.service.go
    |       ├── middleware/
    |       |   └── **.mdlw.go
    ├── schema/
    |      ├── directives/
    |      |   └── **.go
    |      ├── gql/
    |      |   ├── **.graphqls
    |      |   └── schema.graphqls
    |      ├── migration/
    |      |   └── migration.go
    |      ├── model/
    |      |   ├── **.model.go
    |      |   └── models_gen.go
    |      ├── resolver/
    |      |   ├── resolver.go
    |      |   ├── schema.resolver.go
    |      |   └── **.resolver.go
    └── utils/
        └── **.util.go
```

### Description of Directories and Files

- **`gqlgen.yml`**: Configuration file for gqlgen. It defines schema paths, resolver mappings, and code generation settings.

- **`go.mod`**: Go module file that specifies the module's path and its dependencies.

- **`.env`**: Environment variables file for configuring the application (e.g., database credentials, API keys).

- **`server.go`**: Entry point of the application. It initializes configurations, sets up the server, applies middlewares, and starts listening for requests.

- **`src/`**: Contains all the source code organized into subdirectories based on functionality.

  - **`app.handlers.go`**: Handles the main application routes and request handling logic.

  - **`config/`**:

    - **`database.go`**: Handles database connection setup, loading configurations, and initializing database clients.

  - **`generate/`**:

    - **`generate.go`**: Contains scripts or commands related to code generation, such as generating resolvers or models using gqlgen.

  - **`modules/`**: Each module corresponds to a specific feature or domain of the application.

    - **`repository/`**:

      - **`**.repo.go`\*\*: Repository implementations following the repository pattern. These files abstract data access logic, providing methods to interact with the data source (e.g., database queries).

    - **`service/`**:

      - **`**.service.go`\*\*: Service layer implementations containing business logic. Services orchestrate operations, apply business rules, and interact with repositories.

    - **`middleware/`**:
      - **`**.mdlw.go`**: Middleware implementations. The naming convention uses abbreviations like `mdlw` to keep filenames concise. Examples include authentication middleware (`auth.mdlw.go`), logging middleware (`logging.mdlw.go`), etc.

  - **`schema/`**: Contains everything related to GraphQL schemas and resolvers.

    - **`directives/`**:

      - **`**.go`\*\*: Custom directives for GraphQL schema to modify or extend the behavior of GraphQL operations.

    - **`gql/`**:

      - **`**.graphqls`**: Individual GraphQL schema files. Using multiple `.graphqls` files allows for modular schema definitions.
      - **`schema.graphqls`**: The main GraphQL schema file that may import or reference other schema files.

    - **`migration/`**:

      - **`migration.go`**: Handles database migrations, ensuring that the database schema is up to date.

    - **`model/`**:

      - **`**.model.go`\*\*: Defines data models representing entities in the application. Each model corresponds to a database table or a GraphQL type.
      - **`models_gen.go`**: Generated code for models, possibly created by gqlgen or another tool, containing boilerplate code or type definitions.

    - **`resolver/`**:
      - **`resolver.go`**: Sets up the root resolver, connecting the GraphQL schema to resolver implementations.
      - **`schema.resolver.go`**: Generated resolver code that maps schema types and fields to resolver functions.
      - **`**.resolver.go`\*\*: Specific resolver implementations for different parts of the GraphQL schema, handling the logic for fetching and manipulating data.

  - **`utils/`**:
    - **`**.util.go`\*\*: Utility functions and helper methods used across the application, such as error handling, data formatting, etc.

## Setup and Installation

To set up and run this project locally, follow these steps:

1.  **Clone the Repository**:

    ```bash
    git clone https://github.com/yourusername/yourproject.git
    cd yourproject
    ```

2.  **Set Up Environment Variables**:

        - Create a `.env` file in the root directory and populate it with necessary environment variables. Example:

          ```
         ### Database configuration variables ###

          PORT = 3000
          DB_NAME = goweb_db
          DB_USER = root
          DB_PASSWORD = password
          DB_HOST = localhost
          DB_PORT = 3306

        ### Secret application key ###
          JWT_SECRET = hashsecret
          JWT_EXPIRATION_SECRET = 30m
          JWT_REFRESH_SECRET = hashsecret-refresh
          JWT_EXPIRATION_SECRET_REFRESH = 30d
          JWT_EXPIRATION_SECRET_EMAIL = 15m
```

3. **Install Dependencies**:

   Ensure you have Go installed (version 1.16 or higher is recommended). Then run:

   ```bash
   go mod download
   ```

4. **Generate GraphQL Code**:

   Use gqlgen to generate the necessary GraphQL code:

   ```bash
   go run github.com/99designs/gqlgen generate
   ```

   Alternatively, if you have a custom generate script:

   ```bash
   go run src/generate/generate.go
   ```

5. **Build the Application** (Optional):

   To build the application binary:

   ```bash
   go build -o api server.go
   ```

## Usage

To start the API server, run:

```bash
go run server.go
```

Or, if you've built the binary:

```bash
./api
```

The server will start on the configured port (default is `:8080`). You can access the GraphQL playground (if enabled) at `http://localhost:8080` to interact with the API.

## Contributing

Contributions are welcome! To contribute to this project, please follow these steps:

1. **Fork the Repository**: Create a fork of the repository to your GitHub account.

2. **Create a Feature Branch**:

   ```bash
   git checkout -b feature/your-feature-name
   ```

3. **Commit Your Changes**:

   ```bash
   git commit -m "Add feature: your feature description"
   ```

4. **Push to the Branch**:

   ```bash
   git push origin feature/your-feature-name
   ```

5. **Open a Pull Request**: Navigate to the repository on GitHub and open a pull request to merge your changes.

Please ensure that your code follows the project's coding standards and passes all tests.

## License

This project is licensed under the [MIT License](LICENSE). See the [LICENSE](LICENSE) file for details.

---

This README provides a comprehensive overview of your project's architecture, directory structure, and setup instructions. It should serve as a solid foundation for anyone interacting with your API project, whether they're setting it up locally, contributing to it, or simply trying to understand its structure.
