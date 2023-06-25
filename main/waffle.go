package main

import (
	"flag"
	"fmt"
	"github.com/erikbryant/waffle/board"
	"github.com/erikbryant/waffle/pathfinder"
	"github.com/erikbryant/waffle/solver"
	"log"
	"os"
	"runtime/pprof"
)

var (
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func main() {
	fmt.Println("Welcome to waffle!")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	serial := "smkupm.w.nknbeui.e.rgaiey/gyywgw.w.ywygwyw.y.wgwyyg" // 515

	waffle := board.Parse(serial)
	solver := solver.New(waffle)
	solver.Solve()
	solver.Print()

	path := pathfinder.New(solver)
	path.Find()
	path.Print()
}
