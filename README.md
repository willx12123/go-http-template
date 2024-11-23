# go-http-project-template

## FrontEnd

For creating the frontend project, it is recommended to use `vite` with the following command:

```bash
pnpm create vite client --template react-ts
```

The `build.sh` script defaults to using the `client` directory as the frontend project and executes the run script.

## Backend

### Basic Directory Structure

```
- server
  - cmd // some commands like gorm code gen or ent code gen or wire code gen
    - generate
  - internal
    - client // Put third party request here. Like HTTP or RPC.
    - dal
      - db // Database operation put here.
        - query // Default GORM generate ORM here.
        - querygen // If you want to use GORM dynamic sql, put them here. And use them in `cmd/generate/main.go`.
        - db.go // Connect database here.
      - redis
      - es
    - handler // gin handlers
    - service // biz logic
    - middleware // gin middleware
    - pkg
      - config // Env config loader.
      - util // Common utils
      - validator // go-validator instance.
      - logger // Put logger instance here. 
    - srver // Put gin instance here. Set routes and static file handler.
    - types
      - model // Database struct.
      - dto // HTTP request params and response struct.
      - constant
  - scripts
  - main.go
  - .env
  - .env.production
  - .env.development
```

### ORM

To generate template code, utilize **GORM gen**. The script for generation is located at server/scripts/dev/generate_query.sh.

Should there be any changes to `internal/types/model` or `internal/dal/db/querygen`, ensure to synchronize the model in `cmd/generate/generate.go`.

More info please read on [gorm](https://gorm.io/docs/) and [gorm_gen](https://gorm.io/gen/).

### Authenticate

The default method for authentication uses JWT in the `Authorization` header. You have the flexibility to implement OAuth2 or other methods as needed.

#### Auth some api

- Login api: `/api/v1/login`
- Register api: `/api/v1/register`

```go
	{
		apiV1 := api.Group("/v1")

		apiV1.POST("/register", handler.User.Register)
		apiV1.POST("/login", handler.User.Login)

		authRequired := apiV1.Group("/")
		authRequired.Use(middleware.JwtAuth)
		//authRequired.POST("/xx", handler)
	}
```

For APIs that require authentication, include them under `authRequired`.

#### Get CurrentUser

In the handler, use `authenticate.GetCurrentUser(ctx)` to retrieve the user associated with the current token.

If you need to add more information to the JWT or modify its expiration time, update the `TokenClaims` struct in `internal/service/authenticate/jwt.go`. By default, it only includes `UserID`.
