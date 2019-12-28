package solve

import (
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
var counter = 0

func Train() {
	bolt.CreateDB()
	bolt.Bolt.Bucket = &bolt.BboltBucket{Name: "list"}
	bolt.CreateBucket(bolt.Bolt.Bucket)

	tree := Node{}
	buildTree(&tree, 1)
	c := cube.NewRubik()
	runTraining(c, &tree)
	fmt.Println("\nCounter:", counter)
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

func runTraining(c *cube.Rubik, node *Node) {
	addStateToCacheDB(c, node)
	counter++
	if node.Children[0] == nil {
		if counter%100 == 0 {
			fmt.Print("\rTraining: ", counter, " combinations tested.")
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
