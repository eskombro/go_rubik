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
	fmt.Print("   ----------->    ")
	for iterations != 0 {
		r := RandomMove(c, false)
		fmt.Printf("%s ", move_options[r])
		iterations--
	}
	fmt.Println()
}

func RandomMove(c *cube.Rubik, verbose bool) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	ran := r1.Intn(len(move_options))
	move := move_options[ran]
	cube.RotateFace(c, move, verbose)
	return ran
}

func applyMove(c *cube.Rubik, move byte) {
	cube.RotateFace(c, move_options[move], false)
}

func GetOppositeMove(move byte) byte {
	opposite := byte(0)
	optionsCountQuarter := byte((len(move_options) * 2) / 3)
	if move < optionsCountQuarter {
		if move < optionsCountQuarter/2 {
			opposite = move + optionsCountQuarter/2
		} else {
			opposite = move - optionsCountQuarter/2
		}
	} else {
		opposite = move
	}
	return opposite
}

func unapplyMove(c *cube.Rubik, move byte) {
	opposite := GetOppositeMove(move)
	cube.RotateFace(c, move_options[opposite], false)
}
