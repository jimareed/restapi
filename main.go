package main

import (
	"io"
	"os"
	"net/http"
)

func main() {
	port, ok := os.LookupEnv("PORT")

	if !ok {
		port = "8080"
	}

	http.HandleFunc("/fruit", items)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":" + port, nil)
}

func items(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[{\"type\":\"apple\"}, {\"type\":\"orange\"}, {\"type\":\"pear\"}]\n")
}

func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"message\":\"OK\"}\n")
}
