package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHelloWorld (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World!\n")
}

func main() {
    http.HandleFunc("/", handleHelloWorld)

    log.Fatal(http.ListenAndServe(":8080", nil))
}
