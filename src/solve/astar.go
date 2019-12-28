package solve

import (
	"fmt"
	"os"

	"go_rubik/src/cube"
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

func SolveAStar(c *cube.Rubik, usesCache bool) string {
	fmt.Println("Launching A*")
	gd := aStarData{[]*ANode{}, []*ANode{}}
	n := createNode(c, nil, -1)
	gd.openList = append(gd.openList, n)
	rounds := 0

	for {
		if len(gd.openList) > 0 {
			rounds++
			fmt.Printf("\rRounds: %d", rounds)
			current := gd.openList[0]
			gd.openList = removeFromList(current, gd.openList)
			gd.closedList = append(gd.closedList, current)
			isSolution, solution := checkIsSolution(current, usesCache)
			if isSolution {
				fmt.Println("Nodes checked:", rounds)
				return solution
			}
			expandNode(move_options, &gd, current)
		} else {
			fmt.Println("Open list is empty (shouldn't happen)")
			os.Exit(1)
		}
	}
}

func checkIsSolution(current *ANode, usesCache bool) (bool, string) {
	solutionCache := "none"
	solution := ""
	isSolution := false
	if usesCache {
		solutionCache = CheckStateInCacheDB(current.Hash)
		if solutionCache != "none" {
			isSolution = true
		} else if current.Hash == ResultCubeHash {
			isSolution = true
		}
	} else {
		if current.Hash == ResultCubeHash {
			isSolution = true
		}
	}
	if isSolution {

		for current.Parent != nil {
			solution = fmt.Sprint(move_options[current.Move], " ", solution)
			current = current.Parent
		}
		if solutionCache != "none" {
			fmt.Println("\n(Found partial solution on cache)")
			solution += solutionCache
		}
		return true, solution
	}
	return false, solution
}

func createNode(c *cube.Rubik, parent *ANode, move int) *ANode {

	// Create a new node
	node := &ANode{
		Cube:   *c,
		Parent: parent,
		Move:   move,
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

func expandNode(posMoves []string, gd *aStarData, current *ANode) {
	for i := range posMoves {
		// if move is in closedList continue
		new := createNode(&current.Cube, current, i)
		if isNodeInList(new, gd.closedList) != nil {
			continue
		}

		// if move is in openList see if new node has better G
		open_node := isNodeInList(new, gd.openList)
		if open_node != nil && new.G >= open_node.G {
			continue
		}

		// Add new node to open list TODO
		// fmt.Println("Node is being added to open list")
		gd.openList = addToList(new, gd.openList)
	}
}

func isNodeInList(node *ANode, list []*ANode) *ANode {
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
