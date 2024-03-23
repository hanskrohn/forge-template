package directories

type Node struct {
	Value         string
	depth         int
	Children      []*Node
	isDir         bool
}