package tree

type Tree struct {
	Id     string
	Height int
	Left   *Tree
	Right  *Tree
	Root   *Tree
	IsLeaf bool
}
