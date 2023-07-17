package main

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
)

type Game struct {
	serial string
	index  int
}

func main() {
	fmt.Println("Welcome to waffle!")

	serial := Game{"rdnnne e neianii s angety/gwywgy w ywygywy w ygwgwg", 542}

	waffle := board.Parse(serial.serial)
	s := solver.New(waffle)
	if s.Solve() {
		path := pathfinder.New(s)
		path.Find()
		path.Print()
	} else {
		s.Print()
	}
}
