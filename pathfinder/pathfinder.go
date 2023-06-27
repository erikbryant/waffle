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
	row     int
	col     int
	have    rune
	want    rune
	sortKey rune
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

// sort sorts swappable in the fewest swaps (selection sort)
func sort(swappable []Swappable) ([]Swappable, []Swap) {
	swaps := []Swap{}

	// TODO: implement sort

	return swappable, swaps
}

// Find finds a shortest path for swapping tiles to get to the solution
func (p *Path) Find() {
	fmt.Println("Finding shortest path...")

	// Collect all tiles that need to be swapped and what letter they want to be
	swappable := []Swappable{}
	for _, tile := range p.solution.Tiles() {
		if tile.Color == board.Green {
			continue
		}
		want := p.solution.GetSolution(tile.Row, tile.Col)
		swappable = append(swappable, Swappable{tile.Row, tile.Col, tile.Letter, want, 'X'})
	}

	// Sort 'swappable' by 'swappable.want'
	for i := range swappable {
		swappable[i].sortKey = swappable[i].want
	}
	swappable, _ = sort(swappable)

	// Sort 'swappable' by 'swappable.have', recording each swap
	for i := range swappable {
		swappable[i].sortKey = swappable[i].have
	}
	_, p.swaps = sort(swappable)
}

// Print prints a representation of the solver state and shortest path to the console
func (p *Path) Print() {
	p.solution.Print()
	fmt.Println()
	fmt.Printf("A shortest set of swaps is %d swaps:\n", len(p.swaps))
	for _, swap := range p.swaps {
		fmt.Printf("  %c @ (%d, %d) <-> %c @ (%d, %d)\n", swap.l1, swap.r1, swap.c1, swap.l2, swap.r2, swap.c2)
	}
}
