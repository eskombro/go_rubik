package solve

import (
	"encoding/gob"
	"fmt"
	"os"
	"time"

	"go_rubik/src/cube"
)

type CornerSearchNode struct {
	Parent   *CornerSearchNode
	Move     byte
	Depth    byte
	Children []*CornerSearchNode
}

var currentMaxLayers = byte(0)
var startTime time.Time
var totalNodes = 0
var statesMap map[string]byte

const DB_FILE = "cache/corners.bolt"
const DB_BUCKET_CORNERS = "corners"

func CreateCornersTable() {

	// loadSavedData()

	fmt.Println("Known corner states:", len(statesMap))
	startTime = time.Now()

	tree := CornerSearchNode{Move: 100}
	tree.Children = []*CornerSearchNode{}
	statesMap = make(map[string]byte, 1)

	for move := range move_options {
		new := &CornerSearchNode{Depth: 1, Move: byte(move), Parent: nil}
		tree.Children = append(tree.Children, new)
	}

	currentMaxLayers = 0
	for len(statesMap) < 30000000 {
		currentMaxLayers++
		fmt.Println("Launch search with depth:", currentMaxLayers)
		expandNextLayer(&tree, byte(0))
		saveToFile()
		fmt.Println("\n\tSaved", len(statesMap), "states")
	}

	fmt.Println("\nFinal len of map", len(statesMap))
	fmt.Println("States calculated in ", time.Since(startTime))
	fmt.Println("Training finished in ", time.Since(startTime))
}

func loadSavedData() {
	fmt.Println()
	decodeFile, err := os.Open("cache/corners.db")
	if err != nil {
		panic(err)
	}
	defer decodeFile.Close()
	decoder := gob.NewDecoder(decodeFile)
	statesMap = make(map[string]byte)
	err = decoder.Decode(&statesMap)
	if err != nil {
		panic(err)
	}
}

func saveToFile() {
	encodeFile, err := os.Create("cache/corners.db")
	if err != nil {
		panic(err)
	}
	encoder := gob.NewEncoder(encodeFile)
	if err := encoder.Encode(statesMap); err != nil {
		panic(err)
	}
	encodeFile.Close()
}

func expandNextLayer(node *CornerSearchNode, currentDepth byte) {

	if node.Depth == currentMaxLayers-1 && len(node.Children) == 0 {

		// Calculate movements to State
		movesToCurrent := []byte{}
		tmp := node
		movesToCurrent = append([]byte{node.Move}, movesToCurrent...)
		for tmp.Parent != nil {
			movesToCurrent = append([]byte{tmp.Parent.Move}, movesToCurrent...)
			tmp = tmp.Parent
		}

		// Apply movements to State to a new Rubik
		c := cube.NewRubik()
		for _, m := range movesToCurrent {
			cube.RotateFace(c, move_options[m], false)
		}

		// Expand node
		for move := range move_options {
			if !moveIsUseless(byte(move), node.Move) {
				// Apply movement
				applyMove(c, byte(move))
				hash := stateToHash(calculateCornersState(c))
				if statesMap[hash] == 0 {
					new := &CornerSearchNode{Depth: currentDepth + 1, Move: byte(move), Parent: node}
					node.Children = append(node.Children, new)
					statesMap[hash] = node.Depth
					totalNodes++
				}
				unapplyMove(c, byte(move))
				if len(statesMap)%100 == 0 {
					fmt.Printf("\r\tStates calculated: %d | nodes created: %d", len(statesMap), totalNodes)
				}
			}
		}
		// fmt.Println(len(node.Children))
	}
	for _, ch := range node.Children {
		expandNextLayer(ch, currentDepth+1)
	}

}

func moveIsUseless(nextMove byte, lastMove byte) bool {
	isUseless := nextMove == GetOppositeMove(lastMove) ||
		nextMove == lastMove &&
			(nextMove == lastMove+12 || nextMove == lastMove-12) &&
			(nextMove%2 == 1 && lastMove == nextMove-1)
	return isUseless
}

func calculateCornersState(c *cube.Rubik) *[8][3]byte {
	corners := [8][3]byte{
		[3]byte{
			byte(c.Faces[cube.U].Blocks[0]),
			byte(c.Faces[cube.L].Blocks[0]),
			byte(c.Faces[cube.B].Blocks[6]),
		},
		[3]byte{
			byte(c.Faces[cube.U].Blocks[2]),
			byte(c.Faces[cube.B].Blocks[8]),
			byte(c.Faces[cube.R].Blocks[2]),
		},
		[3]byte{
			byte(c.Faces[cube.U].Blocks[6]),
			byte(c.Faces[cube.F].Blocks[0]),
			byte(c.Faces[cube.L].Blocks[2]),
		},
		[3]byte{
			byte(c.Faces[cube.U].Blocks[8]),
			byte(c.Faces[cube.R].Blocks[0]),
			byte(c.Faces[cube.F].Blocks[2]),
		},
		[3]byte{
			byte(c.Faces[cube.D].Blocks[0]),
			byte(c.Faces[cube.L].Blocks[8]),
			byte(c.Faces[cube.F].Blocks[6]),
		},
		[3]byte{
			byte(c.Faces[cube.D].Blocks[2]),
			byte(c.Faces[cube.F].Blocks[8]),
			byte(c.Faces[cube.R].Blocks[6]),
		},
		[3]byte{
			byte(c.Faces[cube.D].Blocks[6]),
			byte(c.Faces[cube.L].Blocks[6]),
			byte(c.Faces[cube.B].Blocks[0]),
		},
		[3]byte{
			byte(c.Faces[cube.D].Blocks[8]),
			byte(c.Faces[cube.R].Blocks[8]),
			byte(c.Faces[cube.B].Blocks[2]),
		},
	}
	return &corners
}

func stateToHash(cornerState *[8][3]byte) string {
	buff := [8]int16{}
	for i, corner := range cornerState {
		buff[i] |= int16(corner[0])
		buff[i] |= int16(corner[1]) << 3
		buff[i] |= int16(corner[2]) << 6
	}
	hash := ""
	for _, tile := range buff {
		hash += fmt.Sprintf("%x", tile)
	}
	return hash
}
