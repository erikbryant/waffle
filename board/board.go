package board

import (
	"fmt"
	"github.com/erikbryant/dictionaries"
	"github.com/fatih/color"
	"math"
	"regexp"
)

const (
	Empty  = ' '
	Border = 'X'
	Green  = 'g'
	Yellow = 'y'
	White  = 'w'
)

// Board implements a widthxheight grid of runes.
type Board struct {
	cells [][]rune
}

// Waffle implements a waffle game board.
type Waffle struct {
	width     int
	height    int
	letters   Board
	colors    Board
	possibles [][][]rune
}

// new creates a new board, populated with empty squares.
func new(width, height int) Board {
	var b Board

	b.cells = make([][]rune, height)
	for row := 0; row < height; row++ {
		b.cells[row] = make([]rune, width)
	}

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			b.cells[row][col] = Empty
		}
	}

	return b
}

// newSlices creates a new board, populated with empty slices.
func newSlices(width, height int) [][][]rune {
	var b [][][]rune

	b = make([][][]rune, height)
	for row := 0; row < height; row++ {
		b[row] = make([][]rune, width)
	}

	return b
}

// New creates a new, empty waffle game.
func New(width, height int) Waffle {
	var w Waffle

	w.width = width
	w.height = height

	w.letters = new(w.width, w.height)
	w.colors = new(w.width, w.height)
	w.possibles = newSlices(w.width, w.height)

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
	return w.letters.cells[row][col], w.colors.cells[row][col]
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
	w.letters.cells[row][col] = l
	w.colors.cells[row][col] = c
}

func (w *Waffle) WhiteTiles() map[rune]int {
	m := map[rune]int{}

	for row := 0; row < w.Height(); row++ {
		for col := 0; col < w.Width(); col++ {
			l, c := w.Get(row, col)
			if c == White {
				m[l]++
			}
		}
	}

	return m
}

func (w *Waffle) YellowDupes() map[rune]int {
	m := map[rune]int{}

	for row := 0; row < w.Height(); row++ {
		for col := 0; col < w.Width(); col++ {
			l, c := w.Get(row, col)
			if c == Yellow {
				m[l]++
			}
		}
	}

	for k, v := range m {
		if v < 2 {
			// TODO: Is this safe?
			delete(m, k)
		}
	}

	return m
}

func (w *Waffle) TilesInRow(row, col int, match rune) map[rune]int {
	m := map[rune]int{}

	// Tiles to the left.
	for colRef := col - 1; colRef >= 0; colRef-- {
		l, c := w.Get(row, colRef)
		if c == Empty || c == Border {
			break
		}
		if c == match {
			m[l]++
		}
	}

	// This tile and ones to the right.
	for colRef := col; colRef < w.Width(); colRef++ {
		l, c := w.Get(row, colRef)
		if c == Empty || c == Border {
			break
		}
		if c == match {
			m[l]++
		}
	}

	return m
}

func (w *Waffle) TilesInCol(row, col int, match rune) map[rune]int {
	m := map[rune]int{}

	// Tiles to the up.
	for rowRef := row - 1; rowRef >= 0; rowRef-- {
		l, c := w.Get(rowRef, col)
		if c == Empty || c == Border {
			break
		}
		if c == match {
			m[l]++
		}
	}

	// This tile and ones to the down.
	for rowRef := row; rowRef < w.Width(); rowRef++ {
		l, c := w.Get(rowRef, col)
		if c == Empty || c == Border {
			break
		}
		if c == match {
			m[l]++
		}
	}

	return m
}

// keys returns a slice of keys from the given map.
func keys(m map[rune]int) []rune {
	p := []rune{}
	for k := range m {
		p = append(p, k)
	}
	return p
}

func (w *Waffle) GetAllLetters() map[rune]int {
	m := map[rune]int{}

	for row := 0; row < w.Height(); row++ {
		for col := 0; col < w.Width(); col++ {
			if row%2 == 1 && col%2 == 1 {
				continue
			}
			l, _ := w.Get(row, col)
			m[l]++
		}
	}

	return m
}

// Print prints a representation of the board state to the console.
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

	for row := 0; row < w.Height(); row++ {
		if row%2 == 1 {
			continue
		}
		re := w.RegexAcross(row)
		fmt.Printf("A%d: egrep '%s' ../dictionaries/wordleGuessable.dict\n", row, re)
	}

	fmt.Println()

	for col := 0; col < w.Width(); col++ {
		if col%2 == 1 {
			continue
		}
		re := w.RegexDown(col)
		fmt.Printf("C%d: egrep '%s' ../dictionaries/wordleGuessable.dict\n", col, re)
	}
}

func parse(serial string) Waffle {
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
