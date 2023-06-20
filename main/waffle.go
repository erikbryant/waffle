package main

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
)

func main() {
	fmt.Println("Welcome to waffle!")

	serial := "smkupm.w.nknbeui.e.rgaiey/gyywgw.w.ywygwyw.y.wgwyyg" // 515

	waffle := board.Parse(serial)
	solver := solver.New(waffle)
	solver.Solve()
	solver.Print()

	path := pathfinder.New(solver)
	path.Find()
	path.Print()
}
