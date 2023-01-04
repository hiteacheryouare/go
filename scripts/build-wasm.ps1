go version;
$Env:GOOS = "js" 
$Env:GOARCH = "wasm"
go build -o .\out\main.wasm .\lib\main.go