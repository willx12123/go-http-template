# go-http-project-template

## FrontEnd

For creating the frontend project, it is recommended to use `vite` with the following command:

```bash
pnpm create vite client --template react-ts
```

The `build.sh` script defaults to using the `client` directory as the frontend project and executes the run script.

## Backend

### ORM

To generate template code, utilize **GORM gen**. The script for generation is located at server/scripts/dev/generate_query.sh.

Should there be any changes to `internal/types/model` or `internal/dal/db/querygen`, ensure to synchronize the model in `cmd/generate/generate.go`.

More info please read on [gorm](https://gorm.io/docs/) and [gorm_gen](https://gorm.io/gen/).

### Authenticate

The default method for authentication uses JWT in the `Authorization` header. You have the flexibility to implement OAuth2 or other methods as needed.

- Login api: `/api/v1/login`
- Register api: `/api/v1/register`
