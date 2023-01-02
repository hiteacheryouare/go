go version;
Set-Variable GOOS=js 
Set-Variable GOARCH=wasm
go build -o .\out\main.wasm .\lib\main.go