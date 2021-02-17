# Build

Build the webassembly binary
`GOOS=js GOARCH=wasm go build -o main.wasm main.go`

Copy the wasm_exec.js helper
`cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`

Serve
`python3 -m http.server`

# Reference

[Compile Go to WebAssembly] (https://fodor.org/blog/go-to-webassembly/)
