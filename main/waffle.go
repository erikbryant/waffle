package main

import (
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/solver"
)

func main() {
	fmt.Println("Welcome to waffle!")

	serial := "tuaehl.r.emrdcnu.i.heoeby/gwgygw.w.wyygwww.g.wgywyg" // 513

	waffle := board.Parse(serial)
	solver := solver.New(waffle)
	solver.Solve()
	solver.Print()
}
