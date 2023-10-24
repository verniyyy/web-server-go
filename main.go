package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Printf("Starting server at port 8080\n")
	http.HandleFunc("/hello", hello)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe("192.168.50.53:8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
	w.WriteHeader(200)
}
