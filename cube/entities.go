package cube

const (
	LINK_TOP = iota
	LINK_BOTTOM
	LINK_LEFT
	LINK_RIGHT
)

const (
	U = iota
	D
	L
	R
	F
	B
)

type Rubik struct {
	Faces [6]Face
}

type Face struct {
	Blocks [9]int
	Links  [4]*Face
}
