package solver

import (
	"fmt"
	"github.com/erikbryant/dictionaries"
	"github.com/erikbryant/waffle/board"
	"golang.org/x/exp/maps"
	"regexp"
	"strings"
)

type Solver struct {
	game      board.Waffle
	possibles [][][]rune
}

// newSlices creates a board.Grid, populated with empty []rune slices
func newSlices(size int) [][][]rune {
	var b [][][]rune

	b = make([][][]rune, size)
	for row := 0; row < size; row++ {
		b[row] = make([][]rune, size)
	}

	return b
}

// New creates an empty waffle game solver
func New(w board.Waffle) Solver {
	var s Solver

	s.game = w
	s.possibles = newSlices(s.Size())

	return s
}

// Size returns the size of the waffle game
func (s *Solver) Size() int {
	return s.game.Size()
}

// Get returns the letter and color at row, col
func (s *Solver) Get(row, col int) (rune, rune) {
	return s.game.Get(row, col)
}

// Set sets the letter and color at row, col
func (s *Solver) Set(row, col int, l, c rune) {
	s.game.Set(row, col, l, c)
}

// setPossibles assigns the set of possible letters to each tile
func (s *Solver) setPossibles() {
	for _, tile := range s.game.Tiles() {
		s.possibles[tile.Row][tile.Col] = s.possibleLetters(tile.Row, tile.Col)
	}
}

// possibleLetters returns the set of all possible letters for the given tile
func (s *Solver) possibleLetters(row, col int) []rune {
	letter, color := s.Get(row, col)

	if color == board.Empty {
		return []rune{}
	}

	if color == board.Green {
		return []rune{letter}
	}

	// The set of possible letters is defined as:
	// pl := w + yd - w(row) - w(col) + y(row) + y(col) - s

	possible := s.WhiteTiles()
	for k := range s.YellowDupes() {
		possible[k]++
	}
	for k := range s.TilesInRow(row, col, board.White) {
		delete(possible, k)
	}
	for k := range s.TilesInCol(row, col, board.White) {
		delete(possible, k)
	}
	for k := range s.TilesInRow(row, col, board.Yellow) {
		possible[k]++
	}
	for k := range s.TilesInCol(row, col, board.Yellow) {
		possible[k]++
	}
	delete(possible, letter)

	return maps.Keys(possible)
}

// WhiteTiles returns the letters on all of the white tiles
func (s *Solver) WhiteTiles() map[rune]int {
	return s.game.Letters(board.White)
}

// YellowDupes returns any yellow tile letters that are duplicates of each other
func (s *Solver) YellowDupes() map[rune]int {
	m := s.game.Letters(board.Yellow)

	for key, val := range m {
		if val < 2 {
			// Remove non-duplicates
			delete(m, key)
		}
	}

	return m
}

// TilesInRow returns the set of letters of a given color (and their count) adjacent to the given coord
func (s *Solver) TilesInRow(row, col int, matchColor rune) map[rune]int {
	m := map[rune]int{}

	// Tiles to the left
	for i := col - 1; i >= 0; i-- {
		l, c := s.Get(row, i)
		if c == board.Empty {
			break
		}
		if c == matchColor {
			m[l]++
		}
	}

	// This tile and ones to the right
	for i := col; i < s.Size(); i++ {
		l, c := s.Get(row, i)
		if c == board.Empty {
			break
		}
		if c == matchColor {
			m[l]++
		}
	}

	return m
}

// TilesInCol returns the set of letters of a given color (and their count) adjacent to the given coord
func (s *Solver) TilesInCol(row, col int, matchColor rune) map[rune]int {
	m := map[rune]int{}

	// Tiles to the up
	for i := row - 1; i >= 0; i-- {
		l, c := s.Get(i, col)
		if c == board.Empty {
			break
		}
		if c == matchColor {
			m[l]++
		}
	}

	// This tile and ones to the down
	for i := row; i < s.Size(); i++ {
		l, c := s.Get(i, col)
		if c == board.Empty {
			break
		}
		if c == matchColor {
			m[l]++
		}
	}

	return m
}

// RegexAcross returns the regular expression of the possible letters for the given row
func (s *Solver) RegexAcross(i int) string {
	if i%2 == 1 {
		return ""
	}

	re := "^"
	for col := 0; col < s.Size(); col++ {
		if len(s.possibles[i][col]) == 1 {
			re += string(s.possibles[i][col][0])
			continue
		}
		re += "["
		for _, l := range s.possibles[i][col] {
			re += string(l)
		}
		re += "]"
	}
	re += "$"

	return re
}

// RegexCown returns the regular expression of the possible letters for the given col
func (s *Solver) RegexDown(i int) string {
	if i%2 == 1 {
		return ""
	}

	re := "^"
	for row := 0; row < s.Size(); row++ {
		if len(s.possibles[row][i]) == 1 {
			re += string(s.possibles[row][i][0])
			continue
		}
		re += "["
		for _, l := range s.possibles[row][i] {
			re += string(l)
		}
		re += "]"
	}
	re += "$"

	return re
}

// YellowEvenRow returns the letters on yellow tiles in non-intersections for the given row
func (s *Solver) YellowEvenRow(i int) []rune {
	y := []rune{}

	for col := 0; col < s.Size(); col++ {
		if col%2 == 0 {
			continue
		}
		l, c := s.Get(i, col)
		if c == board.Yellow {
			y = append(y, l)
		}
	}

	// TODO: return the count of each letter so we can make sure
	// the candidate word isn't using more of them than it should
	return y
}

// YellowEvenCol returns the letters on yellow tiles in non-intersections for the given col
func (s *Solver) YellowEvenCol(i int) []rune {
	y := []rune{}

	for row := 0; row < s.Size(); row++ {
		if row%2 == 0 {
			continue
		}
		l, c := s.Get(row, i)
		if c == board.Yellow {
			y = append(y, l)
		}
	}

	// TODO: return the count of each letter so we can make sure
	// the candidate word isn't using more of them than it should
	return y
}

// MatchWords returns all dictionary words that match the given re and ye criteria
func MatchWords(re string, dict []string, ye []rune) []string {
	matches := []string{}
	for _, word := range dict {
		matched, err := regexp.MatchString(re, word)
		if err != nil {
			fmt.Println("ERROR! 1", err, re, word)
		}
		if matched {
			usesYe := true
			for _, re2 := range ye {
				if !strings.ContainsRune(word, re2) {
					usesYe = false
					break
				}
			}
			if usesYe {
				matches = append(matches, word)
			}
		}
	}
	if len(matches) == 0 {
		fmt.Println("ERROR! 3", re)
	}

	return matches
}

// UniqueLetters returns the letters in a column in a slice of words
func UniqueLetters(words []string, index int) []rune {
	m := map[rune]int{}

	for _, word := range words {
		l := word[index]
		m[rune(l)]++
	}

	return maps.Keys(m)
}

func (s *Solver) GetAllLetters() map[rune]int {
	m := map[rune]int{}

	for _, tile := range s.game.Tiles() {
		m[tile.Letter]++
	}

	delete(m, board.Empty)

	return m
}

// NarrowPossibles finds matches for each word and records those letters
func (s *Solver) narrowPossibles(dict []string) {
	// For each across/down word, look up its regex in dict.
	// For each match, replace possible letters with set
	// of letters from matched words.

	for row := 0; row < s.Size(); row++ {
		if row%2 == 1 {
			continue
		}
		re := s.RegexAcross(row)
		ye := s.YellowEvenRow(row)
		matches := MatchWords(re, dict, ye)

		for col := 0; col < s.Size(); col++ {
			s.possibles[row][col] = UniqueLetters(matches, col)
		}
	}

	for col := 0; col < s.Size(); col++ {
		if col%2 == 1 {
			continue
		}
		re := s.RegexDown(col)
		ye := s.YellowEvenCol(col)
		matches := MatchWords(re, dict, ye)

		for row := 0; row < s.Size(); row++ {
			s.possibles[row][col] = UniqueLetters(matches, row)
		}
	}

	// Now find which letters have an identified final position
	// (the set of possibles is of length one). Subtract these
	// from the list of starting letters. The remainder will be
	// the letters that have yet to be positioned. If any of the
	// possibles (sets > lenght one) contain letters other than
	// these, remove the extraneous letters.

	sl := s.GetAllLetters()

	// Find letters that still need to be placed
	for _, tile := range s.game.Tiles() {
		p := s.possibles[tile.Row][tile.Col]
		if len(p) == 1 {
			// We have narrowed possibles down to just one
			sl[p[0]]--
			if sl[p[0]] == 0 {
				delete(sl, p[0])
			}
		}
	}

	tbp := string(maps.Keys(sl))

	// Remove any letters not in the to-be-placed set
	for _, tile := range s.game.Tiles() {
		p := s.possibles[tile.Row][tile.Col]
		if len(p) > 1 {
			newP := []rune{}
			for _, l := range p {
				matched, err := regexp.MatchString(string(l), tbp)
				if err != nil {
					fmt.Println("ERROR! A", string(l), tbp)
				}
				if matched {
					newP = append(newP, l)
				}
			}
			s.possibles[tile.Row][tile.Col] = newP
		}
	}
}

// Print prints a representation of the solver state to the console
func (s *Solver) Print() {
	s.game.Print()

	for row := 0; row < s.Size(); row++ {
		if row%2 == 1 {
			continue
		}
		re := s.RegexAcross(row)
		fmt.Printf("A%d: egrep '%s' ../dictionaries/merged.dict\n", row, re)
	}

	fmt.Println()

	for col := 0; col < s.Size(); col++ {
		if col%2 == 1 {
			continue
		}
		re := s.RegexDown(col)
		fmt.Printf("C%d: egrep '%s' ../dictionaries/merged.dict\n", col, re)
	}
}

// Solved returns true if the waffle game is solved
func (s *Solver) Solved() bool {
	for _, tile := range s.game.Tiles() {
		p := s.possibles[tile.Row][tile.Col]
		if len(p) > 1 {
			return false
		}
	}
	return true
}

// loadDict returns the guessable word list
func loadDict(wordLen int) []string {
	guessables := dictionaries.LoadFile("../../dictionaries/merged.dict")
	guessables = dictionaries.FilterByLen(guessables, wordLen)
	guessables = dictionaries.SortUnique(guessables)
	return guessables
}

// Solve solves the waffle board game
func (s *Solver) Solve() bool {
	guessables := loadDict(s.Size())

	s.setPossibles()
	attempts := 0
	for !s.Solved() {
		s.narrowPossibles(guessables)
		attempts++
		if attempts > 10 {
			return false
		}
	}

	return true
}
