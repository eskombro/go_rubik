package main

import (
	"fmt"

	"go_rubik/cube"
	"go_rubik/solve"
)

func main() {
	fmt.Println("Rubik's Go!")
	fmt.Println("-----------")

	// solve.Train()

	c := cube.NewRubik()
	solution := solve.Solve(c)
	fmt.Println("----\nSolution: ", solution)
}
