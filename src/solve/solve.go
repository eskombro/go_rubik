package solve

import (
	"crypto/md5"
	"fmt"

	bolt "go_rubik/src/boltdb"
	"go_rubik/src/cube"
)

var ResultCube cube.Rubik
var ResultCubeHash string

func Solve(c *cube.Rubik, useCache bool) string {
	ResultCube := cube.NewRubik()
	ResultCubeHash = GetCubeStateHash(ResultCube)
	openLimit := 2000
	solution := SolveAStar(c, openLimit, useCache)
	return solution
}

func GetCubeStateHash(c *cube.Rubik) string {
	cubeStr := ""
	for _, face := range c.Faces {
		cubeStr += fmt.Sprint(face.Blocks)
	}
	data := []byte(cubeStr)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func CheckStateInCacheDB(hash string) string {
	if bolt.Bolt.Bucket == nil {
		bolt.CreateDB()
		bolt.Bolt.Bucket = &bolt.BboltBucket{Name: "list"}
	}
	solution := bolt.Get(bolt.Bolt.Bucket, hash)
	return solution
}

func addStateToCacheDB(c *cube.Rubik, node *Node) {
	hash := GetCubeStateHash(c)
	solution := ""
	for node.Parent != nil {
		solution += move_options[GetOppositeMove(node.Move)] + " "
		node = node.Parent
	}
	prev := bolt.Get(bolt.Bolt.Bucket, hash)
	if prev == "none" || len(prev) > len(solution) {
		bolt.Put(bolt.Bolt.Bucket, hash, solution)
	}
}
