package main

import (
	"fmt"

	"go_rubik/cube"
)

func main() {
	fmt.Println("Rubik's Go!")
	// Checked U, D, L, R, F, B
	// Checked U', D', L', R', F', B'
	c := cube.NewRubik()
	cube.PrintRubik(c)
	cube.RotateFace(c, "U")
	cube.RotateFace(c, "D")
	cube.RotateFace(c, "L")
	cube.RotateFace(c, "R")
	cube.RotateFace(c, "F")
	cube.RotateFace(c, "B")
	cube.RotateFace(c, "B'")
	cube.RotateFace(c, "U'")
	cube.RotateFace(c, "D'")
	cube.RotateFace(c, "L'")
	cube.RotateFace(c, "R'")
	cube.RotateFace(c, "F'")
	cube.PrintRubik(c)
}
