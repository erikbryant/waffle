package solver

import (
	"fmt"
	"github.com/erikbryant/dictionaries"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/util"
	"regexp"
)

type Solver struct {
	game      board.Waffle
	possibles [][][]rune
}

// newSlices creates a board.Grid, populated with empty []rune slices
func newSlices(width, height int) [][][]rune {
	var b [][][]rune

	b = make([][][]rune, height)
	for row := 0; row < height; row++ {
		b[row] = make([][]rune, width)
	}

	return b
}

// New creates an empty waffle game solver
func New(w board.Waffle) Solver {
	var s Solver

	s.game = w
	s.possibles = newSlices(s.Width(), s.Height())

	return s
}

// Width returns the width of the waffle game
func (s *Solver) Width() int {
	return s.game.Width()
}

// Width returns the height of the waffle game
func (s *Solver) Height() int {
	return s.game.Height()
}

// Get returns the letter and color at row, col
func (s *Solver) Get(row, col int) (rune, rune) {
	return s.game.Get(row, col)
}

// Set sets the letter and color at row, col
func (s *Solver) Set(row, col int, l, c rune) {
	s.game.Set(row, col, l, c)
}

func (s *Solver) GetSolution(row, col int) rune {
	if len(s.possibles[row][col]) != 1 {
		fmt.Println("ERROR! Z3", row, col, s.possibles[row][col])
	}
	return s.possibles[row][col][0]
}

func (s *Solver) SetPossibles() {
	for row := 0; row < s.Height(); row++ {
		for col := 0; col < s.Width(); col++ {
			s.possibles[row][col] = s.PossibleLetters(row, col)
		}
	}
}

// PossibleLetters returns the set of all possible letters for the given tile
func (s *Solver) PossibleLetters(row, col int) []rune {
	letter, color := s.Get(row, col)

	if color == board.Border || color == board.Empty {
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

	return util.Keys(possible)
}

// WhiteTiles returns the letters on all of the white tiles
func (s *Solver) WhiteTiles() map[rune]int {
	m := map[rune]int{}

	for row := 0; row < s.Height(); row++ {
		for col := 0; col < s.Width(); col++ {
			l, c := s.Get(row, col)
			if c == board.White {
				m[l]++
			}
		}
	}

	return m
}

// YellowDupes returns any yellow tile letters that are duplicates of each other
func (s *Solver) YellowDupes() map[rune]int {
	m := map[rune]int{}

	for row := 0; row < s.Height(); row++ {
		for col := 0; col < s.Width(); col++ {
			l, c := s.Get(row, col)
			if c == board.Yellow {
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

// TilesInRow returns the set of letters (and their count) from the given row
func (s *Solver) TilesInRow(row, col int, match rune) map[rune]int {
	m := map[rune]int{}

	// Tiles to the left
	for colRef := col - 1; colRef >= 0; colRef-- {
		l, c := s.Get(row, colRef)
		if c == board.Empty || c == board.Border {
			break
		}
		if c == match {
			m[l]++
		}
	}

	// This tile and ones to the right
	for colRef := col; colRef < s.Width(); colRef++ {
		l, c := s.Get(row, colRef)
		if c == board.Empty || c == board.Border {
			break
		}
		if c == match {
			m[l]++
		}
	}

	return m
}

// TilesInCol returns the set of letters (and their count) from the given col
func (s *Solver) TilesInCol(row, col int, match rune) map[rune]int {
	m := map[rune]int{}

	// Tiles to the up
	for rowRef := row - 1; rowRef >= 0; rowRef-- {
		l, c := s.Get(rowRef, col)
		if c == board.Empty || c == board.Border {
			break
		}
		if c == match {
			m[l]++
		}
	}

	// This tile and ones to the down
	for rowRef := row; rowRef < s.Width(); rowRef++ {
		l, c := s.Get(rowRef, col)
		if c == board.Empty || c == board.Border {
			break
		}
		if c == match {
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
	for col := 0; col < s.Width(); col++ {
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
	for row := 0; row < s.Height(); row++ {
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
func (s *Solver) YellowEvenRow(i int) []string {
	y := []string{}

	for col := 0; col < s.Width(); col++ {
		if col%2 == 0 {
			continue
		}
		l, c := s.Get(i, col)
		if c == board.Yellow {
			y = append(y, "["+string(l)+"]")
		}
	}

	return y
}

// YellowEvenCol returns the letters on yellow tiles in non-intersections for the given col
func (s *Solver) YellowEvenCol(i int) []string {
	y := []string{}

	for row := 0; row < s.Height(); row++ {
		if row%2 == 0 {
			continue
		}
		l, c := s.Get(row, i)
		if c == board.Yellow {
			y = append(y, "["+string(l)+"]")
		}
	}

	return y
}

// MatchWords returns all dictionary words that match the given re and ye criteria
func MatchWords(re string, dict []string, ye []string) []string {
	matches := []string{}
	for _, word := range dict {
		matched, err := regexp.MatchString(re, word)
		if err != nil {
			fmt.Println("ERROR! 1", err, re, word)
		}
		if matched {
			usesYe := true
			for _, re2 := range ye {
				matched, err := regexp.MatchString(re2, word)
				if err != nil {
					fmt.Println("ERROR! 2", err, re, word)
				}
				if !matched {
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

	return util.Keys(m)
}

func (s *Solver) GetAllLetters() map[rune]int {
	m := map[rune]int{}

	for row := 0; row < s.Height(); row++ {
		for col := 0; col < s.Width(); col++ {
			if row%2 == 1 && col%2 == 1 {
				continue
			}
			l, _ := s.Get(row, col)
			m[l]++
		}
	}

	return m
}

// NarrowPossibles finds matches for each word and records those letters
func (s *Solver) NarrowPossibles(dict []string) {
	// For each across/down word, look up its regex in dict.
	// For each match, replace possible letters with set
	// of letters from matched words.

	for row := 0; row < s.Height(); row++ {
		if row%2 == 1 {
			continue
		}
		re := s.RegexAcross(row)
		ye := s.YellowEvenRow(row)
		matches := MatchWords(re, dict, ye)

		for col := 0; col < s.Width(); col++ {
			s.possibles[row][col] = UniqueLetters(matches, col)
		}
	}

	for col := 0; col < s.Width(); col++ {
		if col%2 == 1 {
			continue
		}
		re := s.RegexDown(col)
		ye := s.YellowEvenCol(col)
		matches := MatchWords(re, dict, ye)

		for row := 0; row < s.Height(); row++ {
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
	for row := 0; row < s.Height(); row++ {
		for col := 0; col < s.Width(); col++ {
			if row%2 == 1 && col%2 == 1 {
				continue
			}
			p := s.possibles[row][col]
			if len(p) == 1 {
				sl[p[0]]--
				if sl[p[0]] == 0 {
					delete(sl, p[0])
				}
			}
		}
	}

	tbp := string(util.Keys(sl))

	// Remove any letters not in the to-be-placed set
	for row := 0; row < s.Height(); row++ {
		for col := 0; col < s.Width(); col++ {
			if row%2 == 1 && col%2 == 1 {
				continue
			}
			p := s.possibles[row][col]
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
				s.possibles[row][col] = newP
			}
		}
	}
}

// Print prints a representation of the solver state to the console
func (s *Solver) Print() {
	s.game.Print()

	for row := 0; row < s.Height(); row++ {
		if row%2 == 1 {
			continue
		}
		re := s.RegexAcross(row)
		fmt.Printf("A%d: egrep '%s' ../dictionaries/merged.dict\n", row, re)
	}

	fmt.Println()

	for col := 0; col < s.Width(); col++ {
		if col%2 == 1 {
			continue
		}
		re := s.RegexDown(col)
		fmt.Printf("C%d: egrep '%s' ../dictionaries/merged.dict\n", col, re)
	}
}

// Solved returns true if the waffle game is solved
func (s *Solver) Solved() bool {
	for row := 0; row < s.Height(); row++ {
		for col := 0; col < s.Width(); col++ {
			if row%2 == 1 && col%2 == 1 {
				continue
			}
			p := s.possibles[row][col]
			if len(p) > 1 {
				return false
			}
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
	guessables := loadDict(s.Width())

	s.SetPossibles()
	attempts := 0
	for !s.Solved() {
		s.NarrowPossibles(guessables)
		attempts++
		if attempts > 10 {
			return false
		}
	}

	return true
}
