# Build

Build the webassembly binary
`GOOS=js GOARCH=wasm go build -o ./build/example.wasm main.go`

Copy the wasm_exec.js helper
`cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./build/`

Serve
`python3 -m http.server`

# Reference

[Compile Go to WebAssembly] (https://fodor.org/blog/go-to-webassembly/)
