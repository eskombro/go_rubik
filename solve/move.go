package solve

import (
	"math/rand"
	"time"

	"go_rubik/cube"
)

var move_options = []string{
	"U", "D", "L", "R", "F", "B",
	"U'", "D'", "L'", "R'", "F'", "B'",
	"U2", "D2", "L2", "R2", "F2", "B2",
}

func RandomMove(c *cube.Rubik, verbose bool) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	move := move_options[r1.Intn(len(move_options))]
	cube.RotateFace(c, move, verbose)
}
