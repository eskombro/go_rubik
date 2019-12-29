package solve

import (
	"fmt"
	"time"

	bolt "go_rubik/src/boltdb"
	"go_rubik/src/cube"
)

type Node struct {
	Parent   *Node
	Move     byte
	Depth    int
	Children *[18]*Node
}

var maxDepth = 7
var counter = 0

func Train() {
	bolt.CreateDB("cache/Cache.bolt")
	bolt.CreateBucket("cache/Cache.bolt", "list")

	fmt.Println("Known states:", bolt.CountBucket("cache/Cache.bolt", "list"))
	startTime := time.Now()

	tree := Node{}
	tree.Children = &[18]*Node{}
	buildTree(&tree, 1)
	c := cube.NewRubik()
	runTraining(c, &tree)
	fmt.Println("Training finished in ", time.Since(startTime))
	fmt.Println("\nCombinations counter:", counter)
}

func buildTree(node *Node, depth int) {
	if depth == maxDepth {
		return
	}
	node.Children = &[18]*Node{}
	for move := range move_options {
		node.Children[move] = &Node{Depth: depth, Move: byte(move), Parent: node}
	}
	for _, child := range node.Children {
		buildTree(child, depth+1)
	}
}

func runTraining(c *cube.Rubik, node *Node) {
	addStateToCacheDB("cache/Cache.bolt", "list", c, node)
	counter++
	if node.Children == nil {
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
	if int(node.Move) == len(move_options)-1 {
		node.Parent.Children = nil
	}
}
