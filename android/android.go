package android

import (
	"reaktor/input"
	"strconv"
	"strings"
)

var puzzleInput = "./android/testInput.txt"

// Solve the puzzle
func Solve() string {
	link := NewLink()
	fractions := strings.Split(input.ReadInput(puzzleInput), "\n")
	for _, fraction := range fractions {
		link.readFraction(fraction)
	}
	return link.FindPath()
}

// Coord describes a coordinate pair
type Coord struct {
	x, y int
}

// Link represents a Neural link
type Link struct {
	mapping map[Coord]bool
	start   Coord
	end     map[Coord]bool
}

// NewLink creates a link with the fields instantiated
func NewLink() Link {
	return Link{mapping: make(map[Coord]bool), end: make(map[Coord]bool)}
}

// readFraction adds the path described by a fraction to the neural links mapping.
// If a start/end node is encountered it sets that in the link
func (link *Link) readFraction(fraction string) {
	fractionParts := strings.Split(fraction, " ")
	startCoords := strings.Split(fractionParts[0], ",")
	x, _ := strconv.Atoi(startCoords[0])
	y, _ := strconv.Atoi(startCoords[1])
	coord := Coord{x, y}

	link.mapping[coord] = true
	if len(fractionParts) == 1 {
		return
	}
	for _, dir := range strings.Split(fractionParts[1], ",") {
		switch dir {
		case "D":
			coord.y++
		case "U":
			coord.y--
		case "R":
			coord.x++
		case "L":
			coord.x--
		case "S":
			link.start = coord
		case "F":
			link.end[coord] = true
		}
		link.mapping[coord] = true
	}
}

// FindPath returns a path from the start to a point in the ending of the link
func (link Link) FindPath() string {
	seen := make(map[Coord]bool)
	found, path := link.dfs(link.start, &seen)
	if found {
		return path
	}
	return "Failed to find path"
}

// dfs performs a depth first search from the current coord towards any coordinate in link.end
func (link Link) dfs(currentCoord Coord, seen *map[Coord]bool) (foundEnd bool, path string) {
	if !link.mapping[currentCoord] || (*seen)[currentCoord] {
		return false, ""
	}
	if link.end[currentCoord] {
		return true, ""
	}
	(*seen)[currentCoord] = true
	for _, nextCoord := range currentCoord.adjacent() {
		found, path := link.dfs(nextCoord, seen)
		if found {
			return true, currentCoord.getDirectionTo(nextCoord) + path
		}
	}
	return false, ""
}

// Gets the direction towards an adjacent coordinate.
func (coord Coord) getDirectionTo(nextCoord Coord) string {
	switch {
	case coord.x < nextCoord.x:
		return "R"
	case coord.x > nextCoord.x:
		return "L"
	case coord.y < nextCoord.y:
		return "D"
	case coord.y > nextCoord.y:
		return "U"
	}
	return "FAILED"
}

// returns the coordinates adjacent to the current one (left, right, down or up, but not diagonal).
func (coord Coord) adjacent() []Coord {
	return []Coord{
		{coord.x + 1, coord.y},
		{coord.x - 1, coord.y},
		{coord.x, coord.y + 1},
		{coord.x, coord.y - 1},
	}
}
