package pathfinder

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/solver"
)

// swap contains two tiles to swap
type Swap struct {
	l1 rune
	r1 int
	c1 int
	l2 rune
	r2 int
	c2 int
}

// swappable contains the coordinates of a tile that needs swapping.
type Swappable struct {
	row  int
	col  int
	have rune
	want rune
}

type Path struct {
	solution solver.Solver
	swaps    []Swap
}

// New creates an empty waffle game shortest path finder
func New(s solver.Solver) Path {
	var p Path

	p.solution = s
	p.swaps = []Swap{}

	return p
}

// Size returns the size of the waffle game board
func (p *Path) Size() int {
	return p.solution.Size()
}

// PathLen returns the number of swaps in the path
func (p *Path) PathLen() int {
	return len(p.swaps)
}

// findDouble returns the index of the first double swap (if any)
func findDouble(want, have rune, swappable []Swappable) int {
	for i, swap := range swappable {
		if swap.have == swap.want {
			// The letter is already on the right tile; don't swap it
			continue
		}
		if swap.have == want && swap.want == have {
			return i
		}
	}

	return -999
}

// findFirst returns the index of the first letter that is swappable and matches 'want'
func findFirst(want rune, swappable []Swappable) int {
	for i, swap := range swappable {
		if swap.have == swap.want {
			// The letter is already on the right tile; don't swap it
			continue
		}
		if swap.have == want {
			return i
		}
	}

	// We failed to find a letter to swap; fatal error!
	return -999
}

// sortHave sorts swappable by 'have' in the fewest swaps
func sortHave(swappable []Swappable) ([]Swappable, []Swap) {
	swaps := []Swap{}

	for i := 0; i < len(swappable)-1; i++ {
		if swappable[i].have == swappable[i].want {
			// The letter is already on the right tile; don't swap it
			continue
		}

		// Swap any double swaps first
		index := findDouble(swappable[i].want, swappable[i].have, swappable[i+1:]) + i + 1
		if index < 0 {
			// If no double swap, just swap with the first matching tile
			index = findFirst(swappable[i].want, swappable[i+1:]) + i + 1
		}

		// Record the swap
		s1 := swappable[index]
		s2 := swappable[i]
		swaps = append(swaps, Swap{s1.have, s1.row, s1.col, s2.have, s2.row, s2.col})

		// Swap!
		swappable[i].have, swappable[index].have = swappable[index].have, swappable[i].have
	}

	return swappable, swaps
}

// Find finds a shortest path for swapping tiles to get to the solution
func (p *Path) Find() {
	// Collect all tiles that need to be swapped and what letter they want to be
	swappable := []Swappable{}
	for _, tile := range p.solution.Tiles() {
		if tile.Color == board.Green {
			continue
		}
		want := p.solution.GetSolution(tile.Row, tile.Col)
		swappable = append(swappable, Swappable{tile.Row, tile.Col, tile.Letter, want})
	}

	// Sort 'swappable' by 'swappable.have', recording each swap
	swappable, p.swaps = sortHave(swappable)
}

// Print prints a representation of the solver state and shortest path to the console
func (p *Path) Print() {
	p.solution.Print()
	fmt.Println()
	fmt.Printf("A solution in %d swaps:\n", len(p.swaps))
	for _, swap := range p.swaps {
		fmt.Printf("  %c (%d, %d) <-> %c (%d, %d)\n", swap.l1, swap.r1, swap.c1, swap.l2, swap.r2, swap.c2)
	}
}
