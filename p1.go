package main

import (
	"fmt"
	"net/http"
	"io"
	"time"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}


func main() {
	
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	server := &http.Server{Addr: ":3334", IdleTimeout: time.Duration(30) * time.Second}
	//server := &http.Server{Addr: ":3334"}
	server.ListenAndServe()
}
