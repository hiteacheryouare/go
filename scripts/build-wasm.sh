go version;
export GOOS=js 
export GOARCH=wasm
go build -o ./out/main.wasm ./lib/main.go;