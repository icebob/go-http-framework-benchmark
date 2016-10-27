package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

/*
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}*/

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	addr := "127.0.0.1:8007"
	fmt.Println("Listening on ", addr)

	r := mux.NewRouter()
	/*
		rand.Seed(time.Now().UnixNano())

		for j := 0; j <= 100; j++ {
			path := "/" + RandString(20)
			//fmt.Println("Add route " + path)
			r.HandleFunc(path, handler)
		}*/
	r.HandleFunc("/", handler)
	http.ListenAndServe(addr, r)
}
