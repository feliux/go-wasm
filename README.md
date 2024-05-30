# Web Assembly with Go

Navigating through Go and WebAssembly (WASM) showcases a powerful union, merging Go's concurrent processing with WASM's rapid client-side execution.

# JSON formatter

## Usage

**Copile to wasm**

```bash
$ export PATH="$PATH:$(go env GOPATH)/bin"
$ make GOROOT="$(go env GOROOT)" setup
$ make build
```

**Serve files**

```bash
$ cd cmd/server/
$ go run main.go
# localhost:9090 for a JSON formatter
```

# References

Go to [docs/references.md](./docs/references.md)