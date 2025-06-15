package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func getTiles(w http.ResponseWriter, r *http.Request, landscape []Tile, task []Tile) {
	query := r.PathValue("query")
	regex := QueryToRegexp(query)
	fmt.Fprintf(w, "Processing query: %s\n", query)
	//fmt.Fprintf(w, "The regex is: %v\n", regex)

	var lMatches, tMatches int
	for _, tile := range landscape {
		if tile.Matches(regex){
			fmt.Fprintf(w, "%v\n", tile)
			lMatches++
		}
	}
	fmt.Fprintf(w, "Landscape matches: %v\n", lMatches)

	for _, tile := range task {
		if tile.Matches(regex){
			fmt.Fprintf(w, "%v\n", tile)
			tMatches++
		}
	}
	fmt.Fprintf(w, "Task matches: %v\n", tMatches)
	fmt.Fprintf(w, "Total matches: %v\n", tMatches + lMatches)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("home.html")
	defer file.Close()

	message, _ := io.ReadAll(file)
	fmt.Fprintf(w, string(message))
}

func initTiles(filenames []string) []Tile {
	var tiles []Tile
	for _, filename := range filenames {
		loaded, _ := LoadTilesFromFile("tiles/" + filename + ".txt")
		tiles = append(tiles, loaded...)
	}
	return tiles
}

func main() {
	landscape := initTiles([]string {"landscape", "box_2"})
	task := initTiles([]string{"task", "box_1"})
	getTilesHandler := func(w http.ResponseWriter, r *http.Request) { getTiles(w, r, landscape, task) }
	http.HandleFunc("GET /tiles/{query}", getTilesHandler)

	http.HandleFunc("GET /", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
