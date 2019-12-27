package solve

import (
	"crypto/md5"
	"fmt"
	// "os"

	bolt "go_rubik/boltdb"
	"go_rubik/cube"
)

type Node struct {
	Parent   *Node
	Move     int
	Depth    int
	Children [12]*Node
}

var maxDepth = 6

func Train() {
	bolt.CreateDB()
	bolt.Bolt.Bucket = &bolt.BboltBucket{Name: "list"}
	bolt.CreateBucket(bolt.Bolt.Bucket)

	tree := Node{}
	buildTree(&tree, 1)
	c := cube.NewRubik()
	runTraining(c, &tree)
	fmt.Println("Counter:", counter)
}

func buildTree(node *Node, depth int) {
	if depth == maxDepth {
		return
	}
	for move := range move_options {
		node.Children[move] = &Node{Depth: depth, Move: move, Parent: node}
	}
	for _, child := range node.Children {
		buildTree(child, depth+1)
	}
}

var counter = 0

func runTraining(c *cube.Rubik, node *Node) {
	addStateToDB(c, node)
	counter++
	if node.Children[0] == nil {
		if counter%100 == 0 {
			fmt.Println("Trained:", counter, "solutions found.")
			// 	os.Exit(0)
		}

		unapplyMove(c, node.Move)
		return
	}
	for _, child := range node.Children {
		if child != nil {
			applyMove(c, child.Move)
			runTraining(c, child)
		}
	}
	unapplyMove(c, node.Move)
}

func applyMove(c *cube.Rubik, move int) {
	// fmt.Println("Apply", move)
	cube.RotateFace(c, move_options[move])
}

func GetOppositeMove(move int) int {
	opposite := 0
	optionsCount := len(move_options)
	if move < optionsCount/2 {
		opposite = move + optionsCount/2
	} else {
		opposite = move - optionsCount/2
	}
	return opposite
}

func unapplyMove(c *cube.Rubik, move int) {

	opposite := GetOppositeMove(move)
	// fmt.Println("Unapply", move, "->", opposite)
	cube.RotateFace(c, move_options[opposite])
}

func GetCubeStateHash(c *cube.Rubik) string {
	cubeStr := ""
	for _, face := range c.Faces {
		cubeStr += fmt.Sprint(face.Blocks)
	}
	// fmt.Println("Cube:", cubeStr)
	data := []byte(cubeStr)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func addStateToDB(c *cube.Rubik, node *Node) {
	// Hash state
	hash := GetCubeStateHash(c)
	// Solution to str
	solution := ""
	for node.Parent != nil {
		solution += move_options[GetOppositeMove(node.Move)] + " "
		node = node.Parent
	}
	// Add to DB [todo: Only add if shorter]
	// fmt.Println("====================", hash, solution)
	prev := bolt.Get(bolt.Bolt.Bucket, hash)
	if prev == "none" || len(prev) > len(solution) {
		bolt.Put(bolt.Bolt.Bucket, hash, solution)
	}
}
