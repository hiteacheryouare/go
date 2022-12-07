package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
)

const defaultPort = "3000"

type responseData struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var code int = 200
		w.Header().Set("server", runtime.Version())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		data := responseData{
			Message: "It works!",
			Code:    code,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonData)
	})

	http.HandleFunc("/endpoint", func(w http.ResponseWriter, r *http.Request) {
		// call your function that creates the endpoint here
	})

	fmt.Println("Server is listening at http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
