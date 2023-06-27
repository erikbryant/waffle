package main

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
)

func main() {
	fmt.Println("Welcome to waffle!")

	serial := "smkupm w nknbeui e rgaiey/gyywgw w ywygwyw y wgwyyg" // 515

	waffle := board.Parse(serial)
	s := solver.New(waffle)
	s.Solve()
	path := pathfinder.New(s)
	path.Find()
	path.Print()
}
