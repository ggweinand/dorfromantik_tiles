package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleGetTiles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	http.HandleFunc("/", handleGetTiles)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
