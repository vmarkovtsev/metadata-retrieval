# metadata-retrieval

Current `examples/cmd/` contains an example of how to use the library, implementing a `ghsync` subcmd.

The example cmd can print to sdtout or save to a postgres DB. To help even further with the development, use the options `--log-level=debug --log-http`.

To use, create a personal GitHub token with the scopes **read:org**, **repo**.

```shell
# you can define one or more access tokens (comma separated)
export GITHUB_TOKENS=<xxx>,<yyy>

# Info for individual repositories
go run examples/cmd/*.go repo --version 0 --owner=athenianco --name=metadata-retrieval

# Info for individual organization and its users (not including its repositories)
go run examples/cmd/*.go org --version 0 --name=athenianco

# Info for organization and all its repositories (similar to ghsync deep)
go run examples/cmd/*.go ghsync --version 0 --orgs=athenianco,bblfsh --no-forks
```

To use a postgres DB:

```shell
docker-compose up -d

go run examples/cmd/*.go repo --version 0 --owner=athenianco --name=metadata-retrieval --db=postgres://user:password@127.0.0.1:5432/ghsync?sslmode=disable

docker-compose exec postgres psql postgres://user:password@127.0.0.1:5432/ghsync?sslmode=disable -c "select * from pull_request_reviews"
```

### Migrations

Migrations reside in `database/migrations` and they need to be packed with go-bindata before being usable.
To repack migrations you can use:

```shell
make migration
```

### Testing

To test, run:

```shell
# set your github personal access token (scopes 'read:org', 'repo')
export GITHUB_TOKEN=<xxx>

# start the database if not already running
export POSTGRES_USER=user
export POSTGRES_PASSWORD=password
export POSTGRES_DB=ghsync
docker-compose up -d

# run the tests
export PSQL_USER=${POSTGRES_USER}
export PSQL_PWD=${POSTGRES_PASSWORD}
export PSQL_DB=${POSTGRES_DB}
go test ./...
```

and for coverage information on all the packages, run:

```shell
go test -coverpkg=./... -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```


## Contribute

[Contributions](https://github.com/athenianco/metadata-retrieval/issues) are more than welcome.
