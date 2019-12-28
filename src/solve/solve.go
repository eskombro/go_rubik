package solve

import (
	"fmt"

	bolt "go_rubik/src/boltdb"
	"go_rubik/src/cube"
)

var ResultCube cube.Rubik
var ResultCubeHash string

func mixCube(c *cube.Rubik) {
	fmt.Println("Shuffling Cube:")
	RandomMove(c, true)
	RandomMove(c, true)
	RandomMove(c, true)
	fmt.Println("-----------")
	fmt.Println("Initial state:")
	cube.PrintRubik(c)
	fmt.Println("-----------")
}

func Solve(c *cube.Rubik) string {
	mixCube(c)
	ResultCube := cube.NewRubik()
	ResultCubeHash = GetCubeStateHash(ResultCube)
	solution := SolveAStar(c)
	return solution
}

func CheckStateInCache(c *cube.Rubik) string {
	hash := ""
	if bolt.Bolt.Bucket == nil {
		bolt.CreateDB()
		bolt.Bolt.Bucket = &bolt.BboltBucket{Name: "list"}
	}
	hash = GetCubeStateHash(c)
	solution := bolt.Get(bolt.Bolt.Bucket, hash)
	return solution
}
