package main

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
)

func main() {
	fmt.Println("Welcome to waffle!")

	serial := "tgheea i osreemt a htigrn/gwywgw w wwgggww y wgyyyg" // 522

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
