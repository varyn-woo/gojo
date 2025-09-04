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
1. [Install Go](https://go.dev/doc/install)
2. [`ducksu-protos`](https://github.com/varyn-woo/ducksu-protos)
3. Generate a certificate and key using `openssl` (certificate filepaths must correpond to the ones in the `ListenAndServeTLS` function in `main.go`):
```shell
openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj "/CN=localhost"
```

## Protobufs
If the protobufs are out of sync after the initial submodule setup:
```shell
cd ducksu-protos
git pull
cd ..
buf generate
```

If you made any changes to the protos while working on this repo, make sure to commit your changes and make a pull request in `ducksu-protos`. If you made the most recent changes locally, you only need to generate new protos using:
```shell
buf generate
```

Currently, both Glorious Ducksu and Gojo run locally, so they must be run on the same machine to talk to each other. This will (obviously) be fixed as we work on the project.

## Running the Server
Set up [`glorious-ducksu`](https://github.com/varyn-woo/glorious-ducksu). This way, you can actually run the site and interact with the UI.


To run the server, `cd` into the directory, then:
```shell
go run main.go
```


You may need to visit [`https://localhost:8443/ws`](https://localhost:8443/ws) and force your browser to accept the sus self-signed certificate before the game starts working.