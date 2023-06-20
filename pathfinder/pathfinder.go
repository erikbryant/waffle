package pathfinder

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/solver"
)

// swap contains the coordinates of two tiles to swap.
type swap struct {
	r1 int
	c1 int
	r2 int
	c2 int
}

// swappable contains the coordinates of a tile that needs swapping.
type swappable struct {
  row int
  col int
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

func (p *Path) Width() int {
	return p.solution.Width()
}

func (p *Path) Height() int {
	return p.solution.Height()
}

func (p *Path) RemoveCorrect() {
  // Remove all letters that are already correct
	for row := 0; row < p.Height(); row++ {
		for col := 0; col < p.Width(); col++ {
			if row%2 == 1 && col%2 == 1 {
				continue
			}
			l1, c := p.solution.Get(row, col)
			l2 := p.solution.GetSolution(row, col)
			if l1 == l2 {
				p.solution.Set(row, col, board.Empty, c)
			}
		}
	}
}

func (p *Path) Find() {
	fmt.Println("Finding path...")

  p.RemoveCorrect()

	p.path = append(p.path, swap{0, 1, 0, 2})
}

func (p *Path) Print() {
	p.solution.Print()
	fmt.Println("An optimal path is:", p.path)
}
