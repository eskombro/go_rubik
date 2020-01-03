package main

import (
	"fmt"

	"go_rubik/src/cube"
	"go_rubik/src/solve"
)

func main() {
	// Parameters
	useCache := true
	randomIterations := 9
	trainingSession := false
	createCornersSession := false

	fmt.Println("  .----------------------.")
	fmt.Println("  |      Rubik's Go!     |")
	fmt.Println("  '----------------------'")

	if createCornersSession {
		solve.CreateCornersTable()
	} else if trainingSession {
		solve.Train()
	} else {
		for i := 0; i < 10; i++ {
			fmt.Println("============")
			fmt.Println("= NEW CUBE =")
			c := cube.NewRubik()
			solve.MixCubeRandom(c, randomIterations)
			fmt.Println("------------")
			fmt.Println("Initial state:")
			cube.PrintRubik(c)
			fmt.Println("------------")
			solution := solve.Solve(c, useCache)
			fmt.Println("----\nSolution: ", solution)
			fmt.Println("============")
		}
	}
}
