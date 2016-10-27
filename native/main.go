package main

import (
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	addr := "127.0.0.1:8010"
	fmt.Println("Listening on ", addr)
	http.HandleFunc("/", handler)
	http.ListenAndServe(addr, nil)
}
