package main

import (
	"bufio"
	"os"
	"regexp"
)

type Tile struct {
	layout       string
	doubleLayout string
}

func (t Tile) Matches(r *regexp.Regexp) bool {
	return r.MatchString(t.doubleLayout)
}

func (t Tile) String() string {
	return t.layout
}

func TileFromString(s string) Tile {
	return Tile{s, s + s}
}

func LoadTilesFromFile(filename string) ([]Tile, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tiles []Tile

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tiles = append(tiles, TileFromString(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tiles, nil
}
