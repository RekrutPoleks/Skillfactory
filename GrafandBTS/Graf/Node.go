package Graf

type Node struct {
	key string
	//coast uint

	tops []*Node
}

type Edge struct{
	coast int
	nodeSrc *Node
	nodeDst *Node
}

func NewNode(key string) *Node{
	return &Node{key, make([]*Node,0)}
}