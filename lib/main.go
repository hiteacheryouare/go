package main

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	var root string = "~/souce/repos/go/"
	fileSystem := os.DirFS(root)
	fs.ReadFile(fileSystem, "./pgaes/index.html")
	w.WriteHeader(200)
	w.Header().Add("server", runtime.Version())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hello there, %s!, Welcome to the web.", name)
}
