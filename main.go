package main

import (
	"fmt"

	bolt "go_rubik/boltdb"
	"go_rubik/cube"
	"go_rubik/solve"
)

func main() {
	fmt.Println("Rubik's Go!")
	// Checked U, D, L, R, F, B
	// Checked U', D', L', R', F', B'
	// c := cube.NewRubik()
	// cube.PrintRubik(c)
	// solve.RandomMove(c)
	// cube.PrintRubik(c)

	// solve.Train()

	c := cube.NewRubik()
	solve.RandomMove(c)
	solve.RandomMove(c)
	solve.RandomMove(c)
	solve.RandomMove(c)
	solve.RandomMove(c)
	solve.RandomMove(c)
	solution := "none"
	hash := ""
	if bolt.Bolt.Bucket == nil {
		bolt.CreateDB()
		bolt.Bolt.Bucket = &bolt.BboltBucket{Name: "list"}
	}
	for {
		hash = solve.GetCubeStateHash(c)
		solution = bolt.Get(bolt.Bolt.Bucket, hash)
		if solution != "none" {
			break
		}
		solve.RandomMove(c)
	}
	fmt.Println(solution)
}
