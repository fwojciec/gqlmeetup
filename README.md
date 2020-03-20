# gqlmeetup

This is an example graphql server built using [gqlgen](https://github.com/99designs/gqlgen). It features:

- PostgreSQL backend using [sqlx](https://github.com/jmoiron/sqlx)
- Dataloaders using [dataloaden](https://github.com/vektah/dataloaden)
- Tests
- Session-based authentication and authorization

## Architecture

This project mostly follows Ben Johnson's [standard package layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1). The domain types and interfaces are defined in the root package (`gqlmeetup.go`). The subpackages provide implementations of the interfaces defined by the root package.

### `bcrypt`

This subpackage wraps the standard library bcrypt package and provides methods for checking and hashing user passwords. It provides an implementation of the `PasswordService` interface.

### `cmd`

`cmd` subpackage holds the executables:

- `server` - runs the GraphQL server
- `createuser` - a simple helper tool to create users

### dataloaden

This is where the dataloader code is located. Much of the code is generated using the [dataloaden](https://github.com/vektah/dataloaden) tool. The subpackage provides an implmentation of the `DataLoaderService` interface and it depends on the `Repository` interface for all database-related functionality. The Middleware method is a HTTP middleware which must be used to wrap the GraphQL handler using the dataloader service.

### gqlgen

This subpackage holds the code generated by [gqlgen](https://github.com/99designs/gqlgen) along with custom resolver definitions, directives logic and some helper functions for initializing the handlers used by the HTTP server. This subpackage has the following dependencies:

- `Repository` which provides the resolvers with database functionality;
- `DataLoaderService` which provides the resolvers with dataloader functionality (for optimizing relational queries and avoiding the n+1 query problem);
- `PasswordService` which gives resolvers the ability to check and hash passwords;
- `SessionService` which provides functionality for managing user sessions (login/logout).

### http

This subpackage implements a simple HTTP server that used by the GraphQL server. The package depends on `SessionService` and `DataLoaderService` interface implementations as the server needs to be able to initialize the dataloaders and work with user session on incoming requests.

### scs

This subpackage wraps the [scs](https://github.com/alexedwards/scs) package and provides the functionality related to user authentication and authorization.. It provides an implementation of the `SessionService` interface.

### mocks

Mock versions of the root package interfaces which are used for testing. The mocks are automatically generated using [moq](github.com/matryer/moq).

### postgres

This package provides the database-related functionality to the resolvers, dataloaders and `createuser` executable via its implementation of the `Repository` interface. I'm using the wonderful [sqlx](https://github.com/jmoiron/sqlx) along with some custom helper functions to make working with PostgreSQL easier.
