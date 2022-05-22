# Web Application Boilerplate for Go

This is project is my daily battle tool in building scalable, secure, robust Go web applications. Although maybe not exactly a framework, but the latter word sounds catchier than "project boilerplate", so I took the liberty to call it this way. It has all you need to get started with your project! What does it have out of the box?
* User authentication/authorization
* JWT with Refresh Token logic
* Idiomatic REST endpoints thanks to Fiber framework 
* SQL connectivity via sqlx
* DB migration files
* multi-tier architecture (sometimes called hexagonal architecture)
* DDD in Go
* Propery project structuring for a Go web application
* Use of .env config files
* error management in API

...and more!

I tried to use as few fancy frameworks as possible to make it Go-idiomatic, with the small exception of using Fiber for REST and sqlx for making SQL queries more concise. Hope this doesn't spoil too much the beauty of Go-purism ;) 

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

- Database: Postgresql
- Migration tool: rubenv/sql-migrate 


### .env file
I intentionally left the .env file so that you can start your Database and Authentication without much trouble. The values inside there are purely for local development. Needless to say, you shouldn't ever be sharing .env files in your production ;)
