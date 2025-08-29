# gojo
Go server for handling glorious ducksu backend.

## Dependencies
- [Install Go](https://go.dev/doc/install)
- Install Gin (HTTP server library) -  `go get github.com/gin-gonic/gin`
- Install Cors (add-on middleware to handle CORS preflight HTTP requests) - `go get github.com/gin-contrib/cors`
- Generate a certificate and key using `openssl` (certificate filepaths must correpond to the ones in the `ListenAndServeTLS` function in `main.go`):
```shell
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"
```

## Running the Server
To run the server, `cd` into the directory, then `go run main.go`. You may need to visit [`https://localhost:8443/ws`](https://localhost:8443/ws) and force your browser to accept the sus self-signed certificate before the game starts working.

## Protobufs
See `ducksu-protos/README.md` for more detailed instructions on setting up protos. If you have done all of the setup there, you should just be able to use:
```shell
buf generate
```
to generate the protos for this project.

Currently, both Glorious Ducksu and Gojo run locally, so they must be run on the same machine to talk to each other. This will (obviously) be fixed as we work on the project.
