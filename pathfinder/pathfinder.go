package pathfinder

import (
	"fmt"
	"github.com/erikbryant/waffle/solver"
)

// swap contains two tiles to swap
type swap struct {
	l1 rune
	r1 int
	c1 int
	l2 rune
	r2 int
	c2 int
}

// swappable contains the coordinates of a tile that needs swapping.
type swappable struct {
	row  int
	col  int
	have rune
	want rune
}

type Path struct {
	solution solver.Solver
	path     []swap
}

func New(s solver.Solver) Path {
	var p Path

	p.solution = s

	return p
}

func (p *Path) Size() int {
	return p.solution.Size()
}

func (p *Path) Find() {
	fmt.Println("Finding path...")

	// Do stuff...

	p.path = append(p.path, swap{'a', 0, 1, 'v', 0, 2})
	p.path = append(p.path, swap{'r', 1, 1, 'k', 2, 2})
}

func (p *Path) Print() {
	p.solution.Print()
	fmt.Println()
	fmt.Println("A shortest set of swaps is:")
	for _, p := range p.path {
		fmt.Printf("  %c @ (%d, %d) <-> %c @ (%d, %d)\n", p.l1, p.r1, p.c1, p.l2, p.r2, p.c2)
	}
}
