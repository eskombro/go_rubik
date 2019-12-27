package solve

import (
	// "fmt"

	bolt "go_rubik/boltdb"
	"go_rubik/cube"
)

func mixCube(c *cube.Rubik) {
	RandomMove(c, true)
	RandomMove(c, true)
	RandomMove(c, true)
	RandomMove(c, true)
	RandomMove(c, true)
	// RandomMove(c, true)
	cube.PrintRubik(c)
}

func Solve(c *cube.Rubik) string {
	mixCube(c)
	solution := SolveRandomMethod(c)
	return solution
}

func SolveRandomMethod(c *cube.Rubik) string {
	solution := "none"
	hash := ""
	if bolt.Bolt.Bucket == nil {
		bolt.CreateDB()
		bolt.Bolt.Bucket = &bolt.BboltBucket{Name: "list"}
	}
	for {
		hash = GetCubeStateHash(c)
		solution = bolt.Get(bolt.Bolt.Bucket, hash)
		if solution != "none" {
			break
		}
		RandomMove(c, false)
	}
	return solution
}
