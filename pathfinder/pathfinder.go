package pathfinder

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/solver"
)

// Swap contains two tiles to swap
type Swap struct {
	l1 rune
	r1 int
	c1 int
	l2 rune
	r2 int
	c2 int
}

// Swappable contains the coordinates of a tile that needs swapping
type Swappable struct {
	row  int
	col  int
	have rune
	want rune
}

// Path contains the game and the swaps that solve it
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

// PathLen returns the number of swaps in the final path
func (p *Path) PathLen() int {
	return len(p.swaps)
}

// findTriple returns a single swap that sets up for a double swap
func findTriple(swappable []Swappable) (int, int) {
	for index1 := range swappable {
		for index2 := range swappable {
			if index2 == index1 {
				// Don't swap with ourselves
				continue
			}
			// Are these two tiles swappable?
			if swappable[index1].have == swappable[index2].want {
				for index3 := range swappable {
					if index3 == index1 || index3 == index2 {
						// Don't swap with ourselves
						continue
					}
					if swappable[index2].have == swappable[index3].want &&
						swappable[index1].want == swappable[index3].have {
						return index1, index2
					}
				}
			}
		}
	}

	// No triple swap found
	return -999, -999
}

// findDouble returns the index of the first double swap (if any)
func findDouble(want, have rune, swappable []Swappable) int {
	if want == have {
		fmt.Printf("ERROR: H54 want == have %v %v\n", want, have)
		return -999
	}
	for i, swap := range swappable {
		if swap.have == want && swap.want == have {
			return i
		}
	}

	return -999
}

// findSingle returns the index of the first letter that is swappable and matches 'want'
func findSingle(want rune, swappable []Swappable) int {
	for i, swap := range swappable {
		if swap.have == want {
			return i
		}
	}

	// We failed to find a letter to swap; fatal error!
	return -999
}

// doSwap executes a single swap and removes the swapped tiles from swappable
func doSwap(swappable []Swappable, swaps []Swap, index, i int) ([]Swappable, []Swap) {
	// Construct the swap
	s1 := swappable[index]
	s2 := swappable[i]
	swap := Swap{s1.have, s1.row, s1.col, s2.have, s2.row, s2.col}

	// Record the swap
	swaps = append(swaps, swap)

	// Swap!
	swappable[i].have, swappable[index].have = swappable[index].have, swappable[i].have

	// Remove the swapped tiles
	swappable = removeSwapped(swappable)

	return swappable, swaps
}

// removeSwapped returns the tiles that still need to be swapped
func removeSwapped(swappable []Swappable) []Swappable {
	ns := []Swappable{}

	for i := range swappable {
		if swappable[i].have == swappable[i].want {
			// The letter is already on the right tile
			continue
		}
		ns = append(ns, swappable[i])
	}

	return ns
}

// findPath returns a list of swaps that solve the game
func findPath(swappable []Swappable) []Swap {
	swaps := []Swap{}

	// Keep looping until all swaps are completed
	for len(swappable) > 0 {
		// Do all double swaps
		for i := 0; i < len(swappable); i++ {
			index := findDouble(swappable[i].want, swappable[i].have, swappable[i+1:]) + i + 1
			if index > 0 {
				swappable, swaps = doSwap(swappable, swaps, index, i)
			}
		}

		if len(swappable) == 0 {
			break
		}

		// Is there a swap that will set up for a double swap?
		index1, index2 := findTriple(swappable)
		if index1 > 0 {
			swappable, swaps = doSwap(swappable, swaps, index1, index2)
			// This leaves the double swap to do. Restarting the outer loop will take care of that.
			continue
		}

		// Fall back to a single swap of the first tile
		index := findSingle(swappable[0].want, swappable[1:]) + 1
		swappable, swaps = doSwap(swappable, swaps, index, 0)
	}

	return swaps
}

// thisToThat returns all tiles that need to be swapped and what letter they want to be
func (p *Path) thisToThat() []Swappable {
	swappable := []Swappable{}

	for _, tile := range p.solution.Tiles() {
		if tile.Color == board.Green {
			continue
		}
		want := p.solution.GetSolution(tile.Row, tile.Col)
		swappable = append(swappable, Swappable{tile.Row, tile.Col, tile.Letter, want})
	}

	return swappable
}

// Find finds a shortest path for swapping tiles to get to the solution
func (p *Path) Find() {
	swappable := p.thisToThat()
	p.swaps = findPath(swappable)
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
