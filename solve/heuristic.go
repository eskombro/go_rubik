package solve

import (
	"go_rubik/cube"
)

func CalculateHeuristic(c *cube.Rubik) int {
	return misplacedTiles(c)
}

func misplacedTiles(c *cube.Rubik) int {
	counter := 0
	for f, face := range c.Faces {
		for i := range face.Blocks {
			if face.Blocks[i] != ResultCube.Faces[f].Blocks[i] {
				counter++
			}
		}
	}
	return counter
}
