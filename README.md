# Registration service API on Go

This is an example Go backend service I've created to show my junior colleagues on how to solve ordinary challenges
appearing when you're a writing a back-end service on Go:
* User authentication/authorization
* JWT
* Idiomatic REST endpoints
* DB connection
* DB migrations
* multi-tier architecture (sometimes called hexagonal architecture)
* DDD in Go
* error management in API

...and more!

I tried to use as few fancy frameworks as possible to make it Go-idiomatic, but I couldn't resist using Fiber for REST and sqlx for executing SQL because they both save so much lines of code which were otherwise be wasted in boilerplate code. Hope this doesn't spoil too much the beauty of Go-purism ;) 

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
