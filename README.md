# Wanaku API

## Deployment

1. Set up a PostgresQL server (Docker is the easiest way)
2. Install *rubenv/sql-migrate* as a CLI tool
3. Check if you have the correct DB config set in *./dbconfig.yml*
4. Run `./sql-migrate up` to invoke DB migrations
5. Create an */.env* file in project's root and fill the env variables. The list of env variables necessary to set can
   be found in the main.checkEnvVars() function.
6. Run the app by `./go run main.go`

## Technology stack

### Code stack:

- go/fiber
- sqlx
- godotenv

### DB stack:

- Postgresql
- rubenv/sql-migrate 
