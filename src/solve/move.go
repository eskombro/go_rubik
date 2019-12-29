package solve

import (
	"fmt"
	"math/rand"
	"time"

	"go_rubik/src/cube"
)

var move_options = []string{
	"U", "D", "L", "R", "F", "B",
	"U'", "D'", "L'", "R'", "F'", "B'",
	"U2", "D2", "L2", "R2", "F2", "B2",
}

func MixCubeRandom(c *cube.Rubik, iterations int) {
	fmt.Println("Shuffling Cube:")
	for iterations != 0 {
		RandomMove(c, true)
		iterations--
	}
}

func RandomMove(c *cube.Rubik, verbose bool) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	move := move_options[r1.Intn(len(move_options))]
	cube.RotateFace(c, move, verbose)
}

func applyMove(c *cube.Rubik, move byte) {
	cube.RotateFace(c, move_options[move], false)
}

func GetOppositeMove(move byte) int {
	opposite := 0
	optionsCountQuarter := (len(move_options) * 2) / 3
	if int(move) < optionsCountQuarter {
		if int(move) < optionsCountQuarter/2 {
			opposite = int(move) + optionsCountQuarter/2
		} else {
			opposite = int(move) - optionsCountQuarter/2
		}
	} else {
		opposite = int(move)
	}
	return opposite
}

func unapplyMove(c *cube.Rubik, move byte) {
	opposite := GetOppositeMove(move)
	cube.RotateFace(c, move_options[opposite], false)
}
