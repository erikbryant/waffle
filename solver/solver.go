package solver

import (
	"fmt"
	"github.com/erikbryant/dictionaries"
	"github.com/erikbryant/waffle/board"
	"golang.org/x/exp/maps"
	"regexp"
	"strings"
)

type Possibles [][][]rune

type Solver struct {
	game      board.Waffle
	possibles Possibles
}

// newPossible returns an empty possibles grid
func newPossible(size int) Possibles {
	np := make(Possibles, size)
	for row := 0; row < size; row++ {
		np[row] = make([][]rune, size)
	}
	return np
}

// New creates an empty waffle game solver
func New(w board.Waffle) Solver {
	var s Solver

	s.game = w

	// An empty square of []rune slices
	s.possibles = newPossible(s.Size())

	return s
}

// Size returns the size of the waffle game board
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

// GetSolution returns the solution letter and color at row, col
func (s *Solver) GetSolution(row, col int) rune {
	if len(s.possibles[row][col]) != 1 {
		fmt.Printf("ERROR: length of possibles[%d][%d] is not 1 %v\n", row, col, s.possibles[row][col])
		return board.Empty
	}
	return s.possibles[row][col][0]
}

func (s *Solver) Tiles() []board.Tile {
	return s.game.Tiles()
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

	if row%2 == 1 {
		// This is a standalone tile
		l, c := s.Get(row, col)
		if c == matchColor {
			m[l]++
		}
		return m
	}

	// This is a full word
	for i := 0; i < s.Size(); i++ {
		l, c := s.Get(row, i)
		if c == matchColor {
			m[l]++
		}
	}

	return m
}

// TilesInCol returns the set of letters of a given color (and their count) adjacent to the given coord
func (s *Solver) TilesInCol(row, col int, matchColor rune) map[rune]int {
	m := map[rune]int{}

	if col%2 == 1 {
		// This is a standalone tile
		l, c := s.Get(row, col)
		if c == matchColor {
			m[l]++
		}
		return m
	}

	// This is a full word
	for i := 0; i < s.Size(); i++ {
		l, c := s.Get(i, col)
		if c == matchColor {
			m[l]++
		}
	}

	return m
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

// setPossibles assigns the set of possible letters to each tile
func (s *Solver) setPossibles() {
	for _, tile := range s.game.Tiles() {
		s.possibles[tile.Row][tile.Col] = s.possibleLetters(tile.Row, tile.Col)
	}
}

// regexAcross returns the regular expression of the possible letters for the given row
func (s *Solver) regexAcross(row int) string {
	if row%2 == 1 {
		return ""
	}

	re := "^"
	for i := 0; i < s.Size(); i++ {
		if len(s.possibles[row][i]) == 1 {
			re += string(s.possibles[row][i][0])
			continue
		}
		re += "[" + string(s.possibles[row][i]) + "]"
	}
	re += "$"

	return re
}

// regexDown returns the regular expression of the possible letters for the given col
func (s *Solver) regexDown(col int) string {
	if col%2 == 1 {
		return ""
	}

	re := "^"
	for i := 0; i < s.Size(); i++ {
		if len(s.possibles[i][col]) == 1 {
			re += string(s.possibles[i][col][0])
			continue
		}
		re += "[" + string(s.possibles[i][col]) + "]"
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
	// the candidate word isn't using fewer of them than it should
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
	// the candidate word isn't using fewer of them than it should
	return y
}

// matchWords returns all dictionary words that match the given re and ye criteria
func matchWords(re string, ye []rune, dict []string) []string {
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

// letterUnion returns the set of letters in a given position in slice of words
func letterUnion(words []string, index int) []rune {
	m := map[rune]int{}

	for _, word := range words {
		l := word[index]
		m[rune(l)]++
	}

	return maps.Keys(m)
}

// narrowPossibles finds matches for each word and records those letters
func (s *Solver) narrowPossibles(dict []string) {
	// For each across/down word, look up its regex in dict.
	// For each match, replace possible letters with set
	// of letters from matched words.

	// TODO: Check each match against the set of possible letters.
	// Are there enough possible letters to construct the match
	// word? That is, if the match is m..ch and there is only one
	// 'e' in possible letters then the word 'meech' is not a match.
	// (One too many of the letter 'e'.)

	// For each word across
	for row := 0; row < s.Size(); row++ {
		if row%2 == 1 {
			continue
		}
		re := s.regexAcross(row)
		ye := s.YellowEvenRow(row)
		matches := matchWords(re, ye, dict)

		for col := 0; col < s.Size(); col++ {
			s.possibles[row][col] = letterUnion(matches, col)
		}
	}

	// For each word down
	for col := 0; col < s.Size(); col++ {
		if col%2 == 1 {
			continue
		}
		re := s.regexDown(col)
		ye := s.YellowEvenCol(col)
		matches := matchWords(re, ye, dict)

		for row := 0; row < s.Size(); row++ {
			s.possibles[row][col] = letterUnion(matches, row)
		}
	}

	// Now find which letters have an identified final position
	// (the set of possibles is of length one). Subtract these
	// from the list of starting letters. The remainder will be
	// the letters that have yet to be positioned. If any of the
	// possibles (sets > length one) contain letters other than
	// these, remove the extraneous letters.

	all := s.game.AllLetters()

	// Remove letters that are already placed
	for _, tile := range s.game.Tiles() {
		p := s.possibles[tile.Row][tile.Col]
		if len(p) == 1 {
			// We have narrowed possibles down to just one
			all[p[0]]--
			if all[p[0]] == 0 {
				delete(all, p[0])
			}
		}
	}

	// Remove from possibles letters that are not in to-be-placed
	for _, tile := range s.game.Tiles() {
		p := s.possibles[tile.Row][tile.Col]
		if len(p) > 1 {
			newP := []rune{}
			for _, l := range p {
				if all[l] != 0 {
					newP = append(newP, l)
				}
			}
			s.possibles[tile.Row][tile.Col] = newP
		}
	}
}

// isValidSolution returns whether solution solves the waffle board game
func (s *Solver) isValidSolution(solution Possibles) bool {
	all := s.game.AllLetters()

	for r, row := range solution {
		for c, tile := range row {
			if r%2 == 1 && c%2 == 1 {
				// Skip empty tiles
				continue
			}
			if len(tile) != 1 {
				fmt.Println("ERROR Q3: Expected exactly 1 letter", tile)
				return false
			}
			l := tile[0]
			if all[l] < 1 {
				return false
			}
			all[l]--
		}
	}

	for key, val := range all {
		if val != 0 {
			fmt.Println("ERROR Q4: Expected zero letters", key, val)
			return false
		}
	}

	return true
}

func copy(p Possibles) Possibles {
	newP := newPossible(len(p[0]))

	for r, row := range p {
		for c := range row {
			for _, l := range p[r][c] {
				newP[r][c] = append(newP[r][c], l)
			}
		}
	}

	return newP
}

func (s *Solver) allCombos() []Possibles {
	queue := []Possibles{}

	queue = append(queue, newPossible(s.Size()))

	for r, row := range s.possibles {
		for c, tile := range row {
			if r%2 == 1 && c%2 == 1 {
				// Skip the missing tiles
				continue
			}
			queue2 := []Possibles{}
			for _, entry := range queue {
				for _, l := range tile {
					entry2 := copy(entry)
					entry2[r][c] = []rune{l}
					queue2 = append(queue2, entry2)
				}
			}
			queue = queue2
		}
	}

	return queue
}

func collapseValidPossibles(validPossibles []Possibles) Possibles {
	if len(validPossibles) < 1 {
		fmt.Println("Error Z9: Too few possibles")
		return Possibles{}
	}

	if len(validPossibles) == 1 {
		return validPossibles[0]
	}

	p := newPossible(len(validPossibles[0][0]))

	for row, vpRow := range validPossibles[0] {
		for col := range vpRow {
			if row%2 == 1 && col%2 == 1 {
				// Skip the missing tiles
				continue
			}
			// For each valid possible grid, get the letter at (x, y).
			// Collapse those values into a map. Put those uniqued values
			// into the final grid.
			m := map[rune]int{}
			for _, vp := range validPossibles {
				l := vp[row][col][0]
				m[l]++
			}
			p[row][col] = maps.Keys(m)
		}
	}

	return p
}

// findValidPossibles returns the set of possibles that are compatible with the set of tiles
func (s *Solver) findValidPossibles() {
	vp := []Possibles{}

	for _, c := range s.allCombos() {
		if s.isValidSolution(c) {
			vp = append(vp, c)
		}
	}

	s.possibles = collapseValidPossibles(vp)
}

// Print prints a representation of the solver state to the console
func (s *Solver) Print() {
	s.game.Print()

	for i := 0; i < s.Size(); i++ {
		if i%2 == 1 {
			continue
		}
		fmt.Printf("A%d: egrep '%s' ../dictionaries/waffle.dict\n", i, s.regexAcross(i))
		fmt.Printf("D%d: egrep '%s' ../dictionaries/waffle.dict\n", i, s.regexDown(i))
	}
}

// Solved returns true if the waffle game is solved
func (s *Solver) Solved() bool {
	for _, tile := range s.game.Tiles() {
		p := s.possibles[tile.Row][tile.Col]
		if len(p) != 1 {
			return false
		}
	}
	return true
}

// loadDict returns the guessable word list
func loadDict(wordLen int) []string {
	guessables := dictionaries.LoadFile("../../dictionaries/waffle.dict")
	guessables = dictionaries.FilterByLen(guessables, wordLen)
	guessables = dictionaries.SortUnique(guessables)
	return guessables
}

// Solve solves the waffle board game
func (s *Solver) Solve() bool {
	guessables := loadDict(s.Size())

	count := 0
	s.setPossibles()
	for !s.Solved() {
		count++
		if count > 5 {
			return false
		}
		s.narrowPossibles(guessables)
		s.findValidPossibles()
	}

	return true
}
