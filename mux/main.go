package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	addr := "127.0.0.1:8009"
	fmt.Println("Listening on ", addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	http.ListenAndServe(addr, mux)
}
