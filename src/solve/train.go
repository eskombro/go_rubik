package solve

import (
	"crypto/md5"
	"fmt"

	bolt "go_rubik/src/boltdb"
	"go_rubik/src/cube"
)

type Node struct {
	Parent   *Node
	Move     int
	Depth    int
	Children [18]*Node
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
			fmt.Println("Trained:", counter, "combinations tested.")
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
	if node.Move == len(move_options)-1 {
		node.Parent.Children = [18]*Node{}
	}
}

func applyMove(c *cube.Rubik, move int) {
	cube.RotateFace(c, move_options[move], false)
}

func GetOppositeMove(move int) int {
	opposite := 0
	optionsCountQuarter := (len(move_options) * 2) / 3
	if move < optionsCountQuarter {
		if move < optionsCountQuarter/2 {
			opposite = move + optionsCountQuarter/2
		} else {
			opposite = move - optionsCountQuarter/2
		}
	} else {
		opposite = move
	}
	return opposite
}

func unapplyMove(c *cube.Rubik, move int) {
	opposite := GetOppositeMove(move)
	cube.RotateFace(c, move_options[opposite], false)
}

func GetCubeStateHash(c *cube.Rubik) string {
	cubeStr := ""
	for _, face := range c.Faces {
		cubeStr += fmt.Sprint(face.Blocks)
	}
	data := []byte(cubeStr)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func addStateToDB(c *cube.Rubik, node *Node) {
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
