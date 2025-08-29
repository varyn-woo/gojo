# gojo
Go server for handling glorious ducksu backend.

## Contributing
Now that we have more than one person working on this, some basic contribution rules:
- Do not edit the `main` branch directly (editing `README.md` is an acceptable exception for now but I might make `main` protected... we'll see).
- Create a feature branch with a descriptive name and add your changes there. Remember to rebase regularly to avoid merge conflicts.
- If you create a PR with significant changes:
  - Add the link to your PR to the Jira ticket associated with your change. If there is no associated Jira, make one (you should've made it when starting your task, not when finishing it smh).
    - If Jira integration exists, use that instead of linking. But as of writing this on 8/28/2025, it does not.
  - Test it manually by playing a game to make sure everything still works.
  - Add/edit any appropriate unit tests to confirm the new behavior.
  - If CI tests are up and working, make sure those pass too (ideally it won't even let you merge if they don't).
  - Request a code review from someone else working on the project. If they don't review it in a reasonable amount of time, you can just merge your code at your own peril since this is not a live site yet.
- Once the site goes live, the rules will change! If this `README` is not updated, just ignore this section for now.

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
