package solve

import (
	"go_rubik/src/cube"
)

func CalculateHeuristic(c *cube.Rubik) float64 {
	hash := stateToHash(calculateCornersState(c))
	if statesMap[hash] == 0 {
		return misplacedTiles(c)
	} else {
		return float64(statesMap[hash]) / 8
	}
}

func misplacedTiles(c *cube.Rubik) float64 {
	counter := float64(0)
	for f, face := range c.Faces {
		for i := range face.Blocks {
			if face.Blocks[i] != ResultCube.Faces[f].Blocks[i] {
				counter++
			}
		}
	}
	return counter / 8
}
