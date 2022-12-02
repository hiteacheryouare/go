package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/hello", sayHello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	version := runtime.Version()
	w.WriteHeader(303)
	w.Header().Add("server", version)
	fmt.Fprintf(w, "Welcome to the %s api gateway. To continue, please enter a valid URL", version)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	version := runtime.Version()
	w.Header().Add("server", version)
	fmt.Fprintf(w, "Hello there, %s!, Welcome to the web.", name)
}
