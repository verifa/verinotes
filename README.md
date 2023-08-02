[![Go Report Card](https://goreportcard.com/badge/github.com/verifa/verinotes)](https://goreportcard.com/report/github.com/verifa/verinotes)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
![CI workflow](https://github.com/verifa/verinotes/actions/workflows/ci.yaml/badge.svg)

# VeriNotes

Simple application for note taking, meant for demos.

## Connecting to Postgres

There's a struct that is used with [envconfig](https://github.com/kelseyhightower/envconfig):

```go
// store.go
type Config struct {
	SessionDuration  time.Duration
	PostgresUser     string `split_words:"true"`
	PostgresPassword string `split_words:"true"`
	PostgresDbName   string `default:"verinotes" split_words:"true"`
	PostgresHost     string `split_words:"true"`
	PostgresPort     string `default:"5432" split_words:"true"`
	PostgresSslMode  string `default:"disable" split_words:"true"`
}
```

It's only mandatory to set the host, user and password. If the user is not set VeriNotes will default to using embedded sqlite3.

The prefix for env variables is `VN`, for example:

```bash
export VN_POSTGRES_USER=postgres
export VN_POSTGRES_PASSWORD=veristrongpassword
export VN_POSTGRES_HOST=postgres-postgresql.verinotes.svc.cluster.local
```

Note that the password is printed into logs which is typically fine in demos, maybe it could be hidden behind a flag though...

## Build with ko

You can build VeriNotes without Docker, just use [ko](https://ko.build). There's already a `.ko.yaml` file in the repo. BUT note that you must change to a different base image if you want to produce an image which works with the embedded SQLite since it needs CGO and `ko` by default builds a minimalistic image with `CGO_ENABLE=0`. In future maybe we build multiple tags so there's one that only works with Postgres and one that works standalone.

## Developing

```bash
make ent-gen
make be-dev
make fe-dev
```

Now you can deploy the Go backend and Svelte frontend separately when you change one or the other. Svelte of course hot reloads, but the backend you have to keep rebuilding.

