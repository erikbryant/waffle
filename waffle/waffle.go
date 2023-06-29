package main

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
)

func main() {
	fmt.Println("Welcome to waffle!")

	serial := "ctcrms n drnoaal h rmaihy/gwwwgw y wyygyyw w wgwywg" // 524

	waffle := board.Parse(serial)
	s := solver.New(waffle)
	if s.Solve() {
		path := pathfinder.New(s)
		path.Find()
		path.Print()
	} else {
		s.Print()
	}
}
