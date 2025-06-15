package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getTiles(w http.ResponseWriter, r *http.Request, tiles []Tile) {
	query := r.PathValue("query")
	regex := QueryToRegexp(query)
	fmt.Fprintf(w, "Processing query: %s\n", query)
	fmt.Fprintf(w, "The regex is: %v\n", regex)

	var matches int
	for _, tile := range tiles {

		if tile.Matches(regex){
			fmt.Fprintf(w, "%v\n", tile)
			matches++
		}
	}
	fmt.Fprintf(w, "Total amount of matches: %v\n", matches)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("home.html")
	defer file.Close()

	message, _ := io.ReadAll(file)
	fmt.Fprintf(w, string(message))
}


func main() {
	taskTiles, _ := LoadTilesFromFile("tiles/task_tiles.txt")
	taskTilesHandler := func(w http.ResponseWriter, r *http.Request) { getTiles(w, r, taskTiles) }
	http.HandleFunc("GET /tiles/{query}", taskTilesHandler)

	http.HandleFunc("GET /", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
