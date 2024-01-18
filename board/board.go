package board

import (
	"fmt"
	"math"

	"github.com/fatih/color"
)

const (
	Empty  = ' '
	Green  = 'g'
	Yellow = 'y'
	White  = 'w'
)

type Grid [][]rune

// Waffle implements a waffle game board
type Waffle struct {
	size    int
	letters Grid
	colors  Grid
}

type Tile struct {
	Row    int
	Col    int
	Letter rune
	Color  rune
}

// newGrid creates a new grid, populated with empty squares
func newGrid(size int) Grid {
	grid := make(Grid, size)
	for row := 0; row < size; row++ {
		grid[row] = make([]rune, size)
	}

	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			grid[row][col] = Empty
		}
	}

	return grid
}

// New creates an empty waffle game board
func New(size int) Waffle {
	var w Waffle

	w.size = size
	w.letters = newGrid(w.Size())
	w.colors = newGrid(w.Size())

	return w
}

// Size returns the size of the waffle game
func (w *Waffle) Size() int {
	return w.size
}

// Get returns the letter and color at row, col
func (w *Waffle) Get(row, col int) (rune, rune) {
	if row < 0 || row >= w.Size() || col < 0 || col >= w.Size() {
		return Empty, Empty
	}
	return w.letters[row][col], w.colors[row][col]
}

// Set sets the letter and color at row, col
func (w *Waffle) Set(row, col int, l, c rune) {
	// We are outside the bounds of the grid
	if row < 0 || row >= w.Size() || col < 0 || col >= w.Size() {
		return
	}

	// If row and col are odd, this is a hole in the grid
	if row%2 == 1 && col%2 == 1 {
		return
	}

	w.letters[row][col] = l
	w.colors[row][col] = c
}

// Tiles returns a slice containing every tile on the waffle game board
func (w *Waffle) Tiles() []Tile {
	tiles := []Tile{}

	for row := 0; row < w.Size(); row++ {
		for col := 0; col < w.Size(); col++ {
			l, c := w.Get(row, col)
			if l != Empty {
				tiles = append(tiles, Tile{row, col, l, c})
			}
		}
	}

	return tiles
}

// TilesInRow returns the set of letters of a given color (and their count) adjacent to the given coord
func (w *Waffle) TilesInRow(row, col int, matchColor rune) map[rune]int {
	m := map[rune]int{}

	if row%2 == 1 {
		// This is a standalone tile
		l, c := w.Get(row, col)
		if c == matchColor {
			m[l]++
		}
		return m
	}

	// This is a full word
	for i := 0; i < w.Size(); i++ {
		l, c := w.Get(row, i)
		if c == matchColor {
			m[l]++
		}
	}

	return m
}

// TilesInCol returns the set of letters of a given color (and their count) adjacent to the given coord
func (w *Waffle) TilesInCol(row, col int, matchColor rune) map[rune]int {
	m := map[rune]int{}

	if col%2 == 1 {
		// This is a standalone tile
		l, c := w.Get(row, col)
		if c == matchColor {
			m[l]++
		}
		return m
	}

	// This is a full word
	for i := 0; i < w.Size(); i++ {
		l, c := w.Get(i, col)
		if c == matchColor {
			m[l]++
		}
	}

	return m
}

// Letters returns the letters of a given color and their count
func (w *Waffle) Letters(c rune) map[rune]int {
	m := map[rune]int{}

	for _, tile := range w.Tiles() {
		if tile.Color == c {
			m[tile.Letter]++
		}
	}

	return m
}

// AllLetters returns all the letters and their count
func (w *Waffle) AllLetters() map[rune]int {
	m := map[rune]int{}

	for _, tile := range w.Tiles() {
		m[tile.Letter]++
	}

	return m
}

// maskForColor returns the console text mask for the given tile color
func maskForColor(c rune) *color.Color {
	mask := color.New(color.FgWhite, color.Bold)

	switch c {
	case Green:
		mask = mask.Add(color.BgGreen)
	case Yellow:
		mask = mask.Add(color.BgYellow)
	case White:
		mask = color.New(color.FgBlack)
		mask = mask.Add(color.BgWhite)
	case Empty:
	default:
		mask = mask.Add(color.BgRed)
	}

	return mask
}

// Print prints the waffle game board state to the console
func (w *Waffle) Print() {
	fmt.Printf("Waffle (%dx%d)\n", w.Size(), w.Size())

	for row := 0; row < w.Size(); row++ {
		for col := 0; col < w.Size(); col++ {
			l, c := w.Get(row, col)
			mask := maskForColor(c)
			mask.Printf("%c", l)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}

// Parse unpacks a string into its corresponding waffle game board
func Parse(serial string) Waffle {
	tiles := (len(serial) - 1) / 2
	size := int(math.Sqrt(float64(tiles)))
	w := New(size)

	for row := 0; row < w.Size(); row++ {
		for col := 0; col < w.Size(); col++ {
			l := serial[row*size+col]
			c := serial[row*size+col+(tiles+1)]
			w.Set(row, col, rune(l), rune(c))
		}
	}

	return w
}

// Serialize packs a waffle game board into a string
func (w *Waffle) Serialize() string {
	s := ""

	for row := 0; row < w.Size(); row++ {
		for col := 0; col < w.Size(); col++ {
			l, _ := w.Get(row, col)
			s += string(l)
		}
	}

	s += "/"

	for row := 0; row < w.Size(); row++ {
		for col := 0; col < w.Size(); col++ {
			_, c := w.Get(row, col)
			s += string(c)
		}
	}

	return s
}
