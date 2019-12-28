package main

import (
	"fmt"

	"go_rubik/src/cube"
	"go_rubik/src/solve"
)

func main() {
	// Parameters
	useCache := true
	randomIterations := 8
	trainingSession := false

	fmt.Println("  .----------------------.")
	fmt.Println("  |      Rubik's Go!     |")
	fmt.Println("  '----------------------'")
	c := cube.NewRubik()
	if trainingSession {
		solve.Train()
	} else {
		solve.MixCubeRandom(c, randomIterations)
		fmt.Println("-----------")
		fmt.Println("Initial state:")
		cube.PrintRubik(c)
		fmt.Println("-----------")
		solution := solve.Solve(c, useCache)
		fmt.Println("----\nSolution: ", solution)
	}
}
