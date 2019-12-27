package solve

import (
	"math/rand"
	"time"

	"go_rubik/cube"
)

var move_options = []string{
	"U", "D", "L", "R", "F", "B",
	"U'", "D'", "L'", "R'", "F'", "B'",
}

// var move_options = []string{
// 	"U", "D", "L",
// 	"U'", "D'", "L'",
// }

func RandomMove(c *cube.Rubik) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	move := move_options[r1.Intn(len(move_options))]
	cube.RotateFace(c, move)
}
