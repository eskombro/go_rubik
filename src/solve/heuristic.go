package solve

import (
	// "fmt"
	"go_rubik/src/cube"
	"sort"
)

func CalculateHeuristic(c *cube.Rubik) float64 {
	hash := stateToHash(calculateCornersState(c))
	for i := range CornerTabs {
		if len(CornerTabs[i]) > 1 {
			res := sort.SearchStrings(CornerTabs[i], hash)
			if res < len(CornerTabs[i]) {
				if CornerTabs[i][res] == hash {
					// fmt.Println("FOUND CORNERS")
					return float64(i+1) / 4
				}
			}
		}
	}
	return misplacedTiles(c)
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
	return counter / 4
}
