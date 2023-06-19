package main

import (
	"fmt"
	"github.com/erikbryant/dictionaries"
	"github.com/fatih/color"
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

// PossibleLetters returns the set of all possible letters for the given cell.
func (w *Waffle) PossibleLetters(row, col int) []rune {
	//
	// The possible letters for a cell are given as the set:
	//
	//   pl := g | w + yd - w(row) - w(col) + y(row) + y(col) - ws
	//
	// The match instruction for a word is given as:
	//
	//   mi := [pl1][pl2][pl3][pl4][pl5] | egrep [y2y4]
	//
	// Where:
	//
	// g      = self is green
	// w      = {all white tiles}
	// yd     = {all yellow duplicates}
	// w(row) = {all white tiles in this row}
	// w(col) = {all white tiles in this col}
	// y(row) = {all yellow tiles in this row}
	// y(col) = {all yellow tiles in this col}
	// s      = self
	// y2     = 2nd letter iff yellow
	// y4     = 4th letter iff yellow
	//

	letter, color := w.Get(row, col)

	if color == Border || color == Empty {
		return []rune{}
	}

	if color == Green {
		return []rune{letter}
	}

	//   pl := w + yd - w(row) - w(col) + y(row) + y(col) - s

	possible := w.WhiteTiles()
	for k := range w.YellowDupes() {
		possible[k]++
	}
	for k := range w.TilesInRow(row, col, White) {
		delete(possible, k)
	}
	for k := range w.TilesInCol(row, col, White) {
		delete(possible, k)
	}
	for k := range w.TilesInRow(row, col, Yellow) {
		possible[k]++
	}
	for k := range w.TilesInCol(row, col, Yellow) {
		possible[k]++
	}
	delete(possible, letter)

	return keys(possible)
}

func (w *Waffle) SetPossibles() {
	for row := 0; row < w.Height(); row++ {
		for col := 0; col < w.Width(); col++ {
			w.possibles[row][col] = w.PossibleLetters(row, col)
		}
	}
}

func (w *Waffle) RegexAcross(i int) string {
	if i%2 == 1 {
		return ""
	}

	re := "^"
	for col := 0; col < w.Width(); col++ {
		re += "["
		for _, l := range w.possibles[i][col] {
			re += string(l)
		}
		re += "]"
	}
	re += "$"

	return re
}

func (w *Waffle) RegexDown(i int) string {
	if i%2 == 1 {
		return ""
	}

	re := "^"
	for row := 0; row < w.Height(); row++ {
		re += "["
		for _, l := range w.possibles[row][i] {
			re += string(l)
		}
		re += "]"
	}
	re += "$"

	return re
}

func MatchWords(re string, dict []string) []string {
	matches := []string{}
	for _, word := range dict {
		matched, err := regexp.MatchString(re, word)
		if err != nil {
			fmt.Println("ERROR! 1", err)
		}
		if matched {
			matches = append(matches, word)
		}
	}
	if len(matches) == 0 {
		fmt.Println("ERROR! 2")
	}

	return matches
}

func (w *Waffle) NarrowPossibles(dict []string) {
	// For each across/down word, look up its regex in dict.
	// If no matches, error.
	// For each match, replace possible letters with set
	// of letters from matched words.
	// Repeat until no new information.

	for row := 0; row < w.Height(); row++ {
		if row%2 == 1 {
			continue
		}
		re := w.RegexAcross(row)
		matches := MatchWords(re, dict)

		if len(matches) == 1 {
			for i, ch := range matches[0] {
				w.possibles[row][i] = []rune{ch}
			}
		}
	}

	for col := 0; col < w.Width(); col++ {
		if col%2 == 1 {
			continue
		}
		re := w.RegexDown(col)
		matches := MatchWords(re, dict)

		if len(matches) == 1 {
			for i, ch := range matches[0] {
				w.possibles[i][col] = []rune{ch}
			}
		}
	}
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
	w := New(5, 5)

	for row := 0; row < w.Height(); row++ {
		for col := 0; col < w.Width(); col++ {
			if row%2 == 1 && col%2 == 1 {
				continue
			}
			l := serial[row*5+col]
			c := serial[row*5+col+(5*5+1)]
			w.Set(row, col, rune(l), rune(c))
		}
	}

	return w
}

// loadDict returns the guessable word list
func loadDict(wordLen int) []string {
	guessables := dictionaries.LoadFile("../dictionaries/wordleGuessable.dict")
	guessables = dictionaries.FilterByLen(guessables, wordLen)
	guessables = dictionaries.SortUnique(guessables)
	return guessables
}

func main() {
	fmt.Println("Welcome to waffle!")
	// waffle := parse("eqebla.m.eupirel.n.mdlwal/ggywgw.w.ywygwww.g.wgyywg")
	waffle := parse("tuaehl.r.emrdcnu.i.heoeby/gwgygw.w.wyygwww.g.wgywyg")
	waffle.Print()
	waffle.SetPossibles()
	waffle.Print()
	guessables := loadDict(waffle.Width())
	waffle.NarrowPossibles(guessables)
	waffle.Print()
}
