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
	openLimit := 10000
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

func CheckStateInCacheDB(dbName string, bucketName string, hash string) string {
	if bolt.Bolt[dbName] == nil || bolt.Bolt[dbName].Bucket[bucketName] == nil {
		bolt.CreateDB(dbName)
		bolt.CreateBucket(dbName, "list")
	}
	solution := bolt.Get(dbName, bucketName, hash)
	return solution
}

func addStateToCacheDB(dbName string, bucketName string, c *cube.Rubik, node *Node) {
	hash := GetCubeStateHash(c)
	solution := ""
	for node.Parent != nil {
		solution += move_options[GetOppositeMove(node.Move)] + " "
		node = node.Parent
	}
	prev := bolt.Get(dbName, bucketName, hash)
	if prev == "none" || len(prev) > len(solution) {
		bolt.Put(dbName, bucketName, hash, solution)
	}
}
