## DB
This project using PostgreSQL with [Ent-go](https://entgo.io/) library as ORM.

### New schema
run in `internal` folder
```bash
go run -mod=mod entgo.io/ent/cmd/ent new {{schema_name}}
```
To update schema of db table, go to `internal/ent/schema` and choose your schema.

### Generate
run in `internal` folder
```bash
go generate ./ent
```
*Don't forget to re-generate ent for each time you update schema.*

#### To Generate folder schema for snapshot or intercept
Update generate commang in `ent/generate.go` with:
`//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature intercept,schema/snapshot ./schema`
then, regenerate

### Migration
Auto migration when running go so be careful.

## Swagger OpenAPI
This project using [gin-swagger](https://github.com/swaggo/gin-swagger)
### Swagger UI
```
http://localhost:port/swagger/index.html
```

### Generate swagger
run in `gin-be`
```bash
swag init --parseDependency  --parseInternal -g ./cmd/api/main.go
```
source: https://github.com/swaggo/swag/issues/817

### Documentation
- https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format
- https://github.com/swaggo/swag/blob/master/README.md#api-operation


## Run project
run in `gin-be` folder
```bash
go run ./cmd/api/main.go
```

## Run Test Project
run in `gin-be` folder
```bash
 go test -v ./... 
```


# Project gin-be

One Paragraph of project description goes here

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## MakeFile

run all make commands with clean tests
```bash
make all build
```

build the application
```bash
make build
```

run the application
```bash
make run
```

Create DB container
```bash
make docker-run
```

Shutdown DB container
```bash
make docker-down
```

live reload the application
```bash
make watch
```

run the test suite
```bash
make test
```

clean up binary from the last build
```bash
make clean
```



## Additional

# In Windows, Install gcc for sqllite to work
- https://stackoverflow.com/questions/43580131/exec-gcc-executable-file-not-found-in-path-when-trying-go-build
- https://jmeubank.github.io/tdm-gcc/
