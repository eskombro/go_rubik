package cube

import (
	"fmt"
)

func NewRubik() *Rubik {
	fmt.Println("Creating a Rubik...")
	cube := Rubik{}
	for i := range cube.Faces {
		cube.Faces[i] = NewFace(i)
	}
	cube.Faces[U].Links = [4]*Face{
		&cube.Faces[B],
		&cube.Faces[F],
		&cube.Faces[L],
		&cube.Faces[R],
	}
	cube.Faces[D].Links = [4]*Face{
		&cube.Faces[F],
		&cube.Faces[B],
		&cube.Faces[L],
		&cube.Faces[R],
	}
	cube.Faces[L].Links = [4]*Face{
		&cube.Faces[U],
		&cube.Faces[D],
		&cube.Faces[B],
		&cube.Faces[F],
	}
	cube.Faces[R].Links = [4]*Face{
		&cube.Faces[U],
		&cube.Faces[D],
		&cube.Faces[F],
		&cube.Faces[B],
	}
	cube.Faces[F].Links = [4]*Face{
		&cube.Faces[U],
		&cube.Faces[D],
		&cube.Faces[L],
		&cube.Faces[R],
	}
	cube.Faces[B].Links = [4]*Face{
		&cube.Faces[D],
		&cube.Faces[U],
		&cube.Faces[L],
		&cube.Faces[R],
	}
	return &cube
}

func NewFace(color int) Face {
	face := Face{
		Blocks: [9]int{},
	}
	for i := range face.Blocks {
		face.Blocks[i] = color
	}
	return face
}

func PrintRubik(cube *Rubik) {

	lines := [3]string{}
	strings := []string{
		"\033[0mWH\033[0m",
		"\033[93mYL\033[0m",
		"\033[92mGR\033[0m",
		"\033[94mBL\033[0m",
		"\033[91mRD\033[0m",
		"\033[95mOR\033[0m",
	}
	for _, face := range cube.Faces {
		for _, v := range face.Blocks[:3] {
			lines[0] += fmt.Sprintf("%s ", strings[v])
		}
		for _, v := range face.Blocks[3:6] {
			lines[1] += fmt.Sprintf("%s ", strings[v])
		}
		for _, v := range face.Blocks[6:] {
			lines[2] += fmt.Sprintf("%s ", strings[v])
		}
		lines[0] += "    "
		lines[1] += "    "
		lines[2] += "    "
	}
	fmt.Println(
		"   UP       ",
		"  DOWN      ",
		"  LEFT      ",
		"  RIGHT     ",
		"  FRONT     ",
		"  BACK      ",
	)
	fmt.Println(lines[0])
	fmt.Println(lines[1])
	fmt.Println(lines[2])
}
