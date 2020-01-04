package main

import (
	"fmt"
	"sort"
	"time"

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
		loadData()
		startTime := time.Now()
		cubeNumber := 100
		solvedNumber := 0
		for i := 0; i < cubeNumber; i++ {
			startCubeTime := time.Now()
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
			fmt.Println("Solved in ", time.Since(startCubeTime))
			fmt.Println("============")
		}
		fmt.Println()
		fmt.Printf("===--- SOLVED %d/%d ---===\n", solvedNumber, cubeNumber)
		fmt.Println("Solved in ", time.Since(startTime), "Mean:",
			float64(int(time.Since(startTime).Nanoseconds())/cubeNumber)/1000000000, "s",
		)
	}
}

func loadData() {
	count := 0
	fmt.Println("Loading data")
	if len(solve.CornerTabs[3]) == 0 {
		solve.CornerTabs = solve.LoadCornersSavedData()
		for i := range solve.CornerTabs {
			sort.Strings(solve.CornerTabs[i])
			count += len(solve.CornerTabs[i])
			fmt.Printf("\r%d states loaded", count)
		}
	}
	fmt.Println("\nData loaded")
}
