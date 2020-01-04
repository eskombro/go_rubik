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
		cubeNumber := 100
		solvedNumber := 0
		for i := 0; i < cubeNumber; i++ {
			fmt.Println("============")
			fmt.Printf("= NEW CUBE (%d/%d) =\n", i+1, cubeNumber)
			c := cube.NewRubik()
			solve.MixCubeRandom(c, randomIterations)
			fmt.Println("------------")
			fmt.Println("Initial state:")
			cube.PrintRubik(c)
			fmt.Println("------------")
			solution, ok := solve.Solve(c, useCache)
			if ok {
				solvedNumber++
			}
			fmt.Println("----\nSolution: ", solution)
			fmt.Println("============")
		}
		fmt.Println()
		fmt.Printf("===--- SOLVED %d/%d ---===\n", solvedNumber, cubeNumber)
	}
}
