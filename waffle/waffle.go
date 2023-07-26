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

	serial := Game{"seegni o phuteoo w uednre/gwywgw y wywggyy w ygwwyg", 551}

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
