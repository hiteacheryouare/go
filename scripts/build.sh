go version;
SET GOOS=js 
SET GOARCH=wasm
go build -o ./out/main.wasm ./lib/main.go;