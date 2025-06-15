package tiles

import "regexp"

type Tile struct {
	layout       string
	doubleLayout string
}

func (t Tile) Matches(r regexp.Regexp) bool {
	return r.MatchString(t.doubleLayout)
}

func (t Tile) String() string {
	return t.layout
}

func TileFromString(s string) Tile {
	return Tile{s, s + s}
}
