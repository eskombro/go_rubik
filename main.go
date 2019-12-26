package main

import (
	"fmt"

	"go_rubik/cube"
)

func main() {
	fmt.Println("Rubik's Go!")
	c := cube.NewRubik()
	cube.PrintRubik(c)
	cube.RotateFace(c, "U")
	cube.PrintRubik(c)
	cube.RotateFace(c, "U'")
	cube.PrintRubik(c)
	cube.RotateFace(c, "D")
	cube.PrintRubik(c)
	cube.RotateFace(c, "D'")
	cube.PrintRubik(c)
	cube.RotateFace(c, "L")
	cube.PrintRubik(c)
	cube.RotateFace(c, "L'")
	cube.PrintRubik(c)
	cube.RotateFace(c, "R")
	cube.PrintRubik(c)
	cube.RotateFace(c, "R'")
	cube.PrintRubik(c)
	cube.RotateFace(c, "F")
	cube.PrintRubik(c)
	cube.RotateFace(c, "F'")
	cube.PrintRubik(c)
	cube.RotateFace(c, "B")
	cube.PrintRubik(c)
	cube.RotateFace(c, "B'")
	cube.PrintRubik(c)
}
