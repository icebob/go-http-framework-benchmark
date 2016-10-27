package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	addr := "127.0.0.1:8008"
	fmt.Println("Listening on ", addr)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello world!")
	})

	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Millisecond)
		io.WriteString(w, "Done")
	})

	mux.HandleFunc("/action/", func(w http.ResponseWriter, r *http.Request) {
		// TODO
	})

	http.ListenAndServe(addr, mux)
}
