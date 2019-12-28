package solve

import (
	"fmt"
	"os"

	"go_rubik/cube"
)

type ANode struct {
	Parent   *ANode
	F, G, H  int
	Move     int
	Cube     cube.Rubik
	Hash     string
	Children [18]*ANode
}

type aStarData struct {
	openList   []*ANode
	closedList []*ANode
}

func SolveAStar(c *cube.Rubik) string {
	solution := "none"
	fmt.Println("Launching A*")
	gd := aStarData{[]*ANode{}, []*ANode{}}
	fmt.Println("Result Hash", ResultCubeHash)
	n := createNode(c, nil, -1)
	gd.openList = append(gd.openList, n)
	running := true
	rounds := 0

	for running {
		if len(gd.openList) > 0 {
			rounds++
			fmt.Printf("\rRounds: %d", rounds)
			current := gd.openList[0]
			gd.openList = removeFromList(current, gd.openList)
			gd.closedList = append(gd.closedList, current)
			if current.Hash == ResultCubeHash {
				fmt.Println("\rFOUND SOLUTION")
				fmt.Println("Nodes checked:", rounds)
				os.Exit(0)
			}
			handleNode(move_options, &gd, current)
		} else {
			fmt.Println("Open list is empty (shouldn't happen)")
			fmt.Println("Did", rounds, "rounds")
			os.Exit(1)
		}
	}
	return solution
}

func createNode(c *cube.Rubik, parent *ANode, move int) *ANode {

	// Create a new node
	node := &ANode{
		Cube:   *c,
		Parent: parent,
	}
	cube.UpdateCubeLinks(&node.Cube)
	// Apply movement and calculate hash
	if move != -1 {
		cube.RotateFace(&node.Cube, move_options[move], false)
	}
	node.Hash = GetCubeStateHash(&node.Cube)
	//Calculate Heuristic
	if parent != nil {
		node.G = parent.G + 1
	}
	node.H = CalculateHeuristic(&node.Cube)
	node.F = node.G + node.H
	return node
}

func handleNode(posMoves []string, gd *aStarData, current *ANode) {
	for i := range posMoves {
		// if move is in closedList continue
		new := createNode(&current.Cube, current, i)
		if tabInSlice(new, gd.closedList) != nil {
			continue
		}

		// if move is in openList see if new node has better G
		open_node := tabInSlice(new, gd.openList)
		if open_node != nil && new.G >= open_node.G {
			continue
		}

		// Add new node to open list TODO
		// fmt.Println("Node is being added to open list")
		gd.openList = addToList(new, gd.openList)
	}
}

func tabInSlice(node *ANode, list []*ANode) *ANode {
	for _, b := range list {
		if node.Hash == b.Hash {
			return b
		}
	}
	return nil
}

func addToList(new *ANode, list []*ANode) []*ANode {
	list = append(list, new)
	for i, n := range list {
		if new.F <= n.F {
			// COULD BE ORDERED BY LOWEST G AS SECND CRITERIA
			copy(list[i+1:], list[i:])
			list[i] = new
			break
		}
	}
	return (list)
}

func removeFromList(todelete *ANode, list []*ANode) []*ANode {
	tmp := 0
	for i := range list {
		if list[i] == todelete {
			tmp = i
			break
		}
	}
	return append(list[:tmp], list[tmp+1:]...)
}
