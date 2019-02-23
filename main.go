package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/fruit", items)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}

func items(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[{\"type\":\"apple\"}, {\"type\":\"orange\"}, {\"type\":\"pear\"}]\n")
}

func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"message\":\"OK\"}\n")
}
