package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

var client *rpc.Client

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	// Synchronous call
	args := &Args{rand.Intn(100), rand.Intn(100)}
	var reply int
	err := client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Fprintf(w, "Arith: %d*%d=%d", args.A, args.B, reply)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var err error
	client, err = rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	http.HandleFunc("/hello", handler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(":88", nil)
}
