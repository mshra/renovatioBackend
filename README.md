## Enable Live Reloading for development

1. Install [Air](https://github.com/air-verse/air)

```bash
$ go install github.com/air-verse/air@latest
```

2. Add the following alias to your .bashrc or .zshrc

```bash
alias air='$(go env GOPATH)/bin/air'
```

3. Start the dev server with live reloading enabled

```bash
$ air -c .air.toml
```

4. This will start a dev server with live reloading enabled on [http://localhost:8000/](http://localhost:8000/)

## Running tests

```bash
$ go test -v ./tests
```
