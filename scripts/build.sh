go version;
GOOS=js 
GOARCH=wasm
go build ./lib/main.go -o ./out/main.wasm;