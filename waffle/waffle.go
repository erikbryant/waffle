package main

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
)

func main() {
	fmt.Println("Welcome to waffle!")

	serial := "caceki e lltsaoe i nmhlnl/gwywgy y yywgyww w ygwwwg" // 531

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
