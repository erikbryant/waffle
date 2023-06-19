package board

import (
	"fmt"
	"github.com/fatih/color"
	"math"
)

const (
	Empty  = ' '
	Border = 'X'
	Green  = 'g'
	Yellow = 'y'
	White  = 'w'
)

// Waffle implements a waffle game board.
type Waffle struct {
	width   int
	height  int
	letters [][]rune
	colors  [][]rune
}

// new creates a new board, populated with empty squares.
func new(width, height int) [][]rune {
	var board [][]rune

	board = make([][]rune, height)
	for row := 0; row < height; row++ {
		board[row] = make([]rune, width)
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			board[row][col] = Empty
		}
	}

	return board
}

// New creates a new, empty waffle game.
func New(width, height int) Waffle {
	var w Waffle

	w.width = width
	w.height = height

	w.letters = new(w.width, w.height)
	w.colors = new(w.width, w.height)

	return w
}

// Width returns the width of the waffle game.
func (w *Waffle) Width() int {
	return w.width
}

// Height returns the height of the waffle game.
func (w *Waffle) Height() int {
	return w.height
}

// Get returns the letter and color at row,col.
func (w *Waffle) Get(row, col int) (rune, rune) {
	if row < 0 || row >= w.Height() || col < 0 || col >= w.Width() {
		return Border, Border
	}
	// If row and col are odd, this is a gap. Return 'empty'.
	if row%2 == 1 && col%2 == 1 {
		return Empty, Empty
	}
	return w.letters[row][col], w.colors[row][col]
}

// Set sets the letter and color at row,col.
func (w *Waffle) Set(row, col int, l, c rune) {
	if row < 0 || row >= w.Height() || col < 0 || col >= w.Width() {
		return
	}
	// If row and col are odd, this is a gap.
	if row%2 == 1 && col%2 == 1 {
		return
	}
	w.letters[row][col] = l
	w.colors[row][col] = c
}

// Print prints a representation of the [][]rune state to the console.
func (w *Waffle) Print() {
	fmt.Printf("Waffle (%dx%d)\n", w.Width(), w.Height())

	for row := 0; row < w.Height(); row++ {
		for col := 0; col < w.Width(); col++ {
			l, c := w.Get(row, col)
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
			mask.Printf("%c", l)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
}

func Parse(serial string) Waffle {
	tiles := (len(serial) - 1) / 2
	size := int(math.Sqrt(float64(tiles)))
	w := New(size, size)

	for row := 0; row < w.Height(); row++ {
		for col := 0; col < w.Width(); col++ {
			if row%2 == 1 && col%2 == 1 {
				continue
			}
			l := serial[row*size+col]
			c := serial[row*size+col+(tiles+1)]
			w.Set(row, col, rune(l), rune(c))
		}
	}

	return w
}
