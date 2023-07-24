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

	serial := Game{"negeapar e o rdiippdda f p elmterdlu r a teadrlte/ywgwgwww y w ygwgggwgy w y ygwgggwgy y y yywgygwy", 61}

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
