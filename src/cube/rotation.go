package cube

import (
	"fmt"
	"strings"
)

func RotateFace(cube *Rubik, instruction string, verbose bool) {
	instructions := "UDLRFB"
	direction := strings.Index(instructions, string(instruction[0]))
	prime := false
	double := false
	if len(instruction) > 1 {
		prime = string(instruction[1]) == "'"
		double = string(instruction[1]) == "2"
	}
	if verbose {
		fmt.Println("   Rotating cube in direction: ", instruction)
	}
	face := &cube.Faces[direction]
	simpleFaceRotation(face, prime)
	if double {
		simpleFaceRotation(face, prime)
	}
	edgesRotation(face, direction, prime)
	if double {
		edgesRotation(face, direction, prime)
	}
}

func simpleFaceRotation(face *Face, prime bool) {
	if !prime {
		face.Blocks = [9]int{
			face.Blocks[6],
			face.Blocks[3],
			face.Blocks[0],
			face.Blocks[7],
			face.Blocks[4],
			face.Blocks[1],
			face.Blocks[8],
			face.Blocks[5],
			face.Blocks[2],
		}
	} else {
		face.Blocks = [9]int{
			face.Blocks[2],
			face.Blocks[5],
			face.Blocks[8],
			face.Blocks[1],
			face.Blocks[4],
			face.Blocks[7],
			face.Blocks[0],
			face.Blocks[3],
			face.Blocks[6],
		}
	}
}

func edgesRotation(face *Face, direction int, prime bool) {
	var pos_1 [3]int
	var pos_2 [3]int
	var pos_3 [3]int
	var pos_4 [3]int
	if direction == U {
		pos_1 = [3]int{6, 7, 8}
		pos_2 = [3]int{2, 1, 0}
		pos_3 = [3]int{2, 1, 0}
		pos_4 = [3]int{2, 1, 0}
	} else if direction == D {
		pos_1 = [3]int{6, 7, 8}
		pos_2 = [3]int{6, 7, 8}
		pos_3 = [3]int{2, 1, 0}
		pos_4 = [3]int{6, 7, 8}
	} else if direction == L {
		pos_1 = [3]int{0, 3, 6}
		pos_2 = [3]int{0, 3, 6}
		pos_3 = [3]int{0, 3, 6}
		pos_4 = [3]int{0, 3, 6}
	} else if direction == R {
		pos_1 = [3]int{2, 5, 8}
		pos_2 = [3]int{2, 5, 8}
		pos_3 = [3]int{2, 5, 8}
		pos_4 = [3]int{2, 5, 8}
	} else if direction == F {
		pos_1 = [3]int{6, 7, 8}
		pos_2 = [3]int{8, 5, 2}
		pos_3 = [3]int{2, 1, 0}
		pos_4 = [3]int{0, 3, 6}
	} else if direction == B {
		pos_1 = [3]int{8, 7, 6}
		pos_2 = [3]int{6, 3, 0}
		pos_3 = [3]int{0, 1, 2}
		pos_4 = [3]int{2, 5, 8}
	}
	tmp := [3]int{
		face.Links[LINK_TOP].Blocks[pos_1[0]],
		face.Links[LINK_TOP].Blocks[pos_1[1]],
		face.Links[LINK_TOP].Blocks[pos_1[2]],
	}
	for i := 0; i < 3; i++ {
		if !prime {
			face.Links[LINK_TOP].Blocks[pos_1[i]] = face.Links[LINK_LEFT].Blocks[pos_2[i]]
			face.Links[LINK_LEFT].Blocks[pos_2[i]] = face.Links[LINK_BOTTOM].Blocks[pos_3[i]]
			face.Links[LINK_BOTTOM].Blocks[pos_3[i]] = face.Links[LINK_RIGHT].Blocks[pos_4[i]]
			face.Links[LINK_RIGHT].Blocks[pos_4[i]] = tmp[i]
		} else {
			face.Links[LINK_TOP].Blocks[pos_1[i]] = face.Links[LINK_RIGHT].Blocks[pos_4[i]]
			face.Links[LINK_RIGHT].Blocks[pos_4[i]] = face.Links[LINK_BOTTOM].Blocks[pos_3[i]]
			face.Links[LINK_BOTTOM].Blocks[pos_3[i]] = face.Links[LINK_LEFT].Blocks[pos_2[i]]
			face.Links[LINK_LEFT].Blocks[pos_2[i]] = tmp[i]
		}
	}
}
