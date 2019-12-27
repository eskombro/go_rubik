package main

import (
	"fmt"

	"go_rubik/cube"
	"go_rubik/solve"
)

func main() {
	fmt.Println("Rubik's Go!")
	// Checked U, D, L, R, F, B
	// Checked U', D', L', R', F', B'

	// solve.Train()

	c := cube.NewRubik()
	// solve.RandomMove(c, true)
	// solve.RandomMove(c, true)
	// solve.RandomMove(c, true)
	// solve.RandomMove(c, true)
	// solve.RandomMove(c, true)
	cube.PrintRubik(c)
	solution := solve.Solve(c)
	fmt.Println("----\nSolution: ", solution)
}
