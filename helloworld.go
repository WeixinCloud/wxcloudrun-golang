package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	target := "Hello world!"
	fmt.Fprintf(w, "Hello, %s!\n", target)
}

func main() {
	log.Print("Hello world sample started.")
	http.HandleFunc("/", handler)


	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", 80), nil))
}
