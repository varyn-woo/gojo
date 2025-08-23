# gojo
Go server for handling glorious ducksu backend.

## Dependencies
- [Install Go](https://go.dev/doc/install)
- Install Gin (HTTP server library) -  `go get github.com/gin-gonic/gin`
- Install Cors (add-on middleware to handle CORS preflight HTTP requests) - `go get github.com/gin-contrib/cors`

## Running the Server
To run the server, `cd` into the directory, then `go run main.go`.

## Protobufs
See `ducksu-protos/README.md` for more detailed instructions on setting up protos. If you have done all of the setup there, you should just be able to use:
```shell
buf generate
```
to generate the protos for this project.

Currently, both Glorious Ducksu and Gojo run locally, so they must be run on the same machine to talk to each other. This will (obviously) be fixed as we work on the project.
