# Golang WASM drawing fractal trees

Draws random fractal trees on html canvas.
Written in Golang WebAssembly.

Demo http://fractaltree-wasm.surge.sh/

Build command
```sh
$ SET GOOS=js SET GOARCH=wasm go build -o html/main.wasm main.go
```

To test locally
```sh
$ npm i -g http-server
$ cd html
$ http-server -c-1 -o
```
