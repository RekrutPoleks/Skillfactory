package BTS

import "fmt"

type ActionOnTree interface {
	Push(int)error
	Pop(int)(*Node, error)
	Search(int)(*Node, error)
	PrintTree()
}

type Tree struct {
	Node *Node

}

func (t *Tree) Push(key int) error{
	if t.Node == nil {
		t.Node = createNode(key)
		return nil
	}
	var inNode = t.Node
	for {
		switch {
		case key == (*inNode).Key:
			return fmt.Errorf("%d the key exists ", key)
		case (*inNode).Key > key:
			if (*inNode).right == nil{
				(*inNode).right = createNode(key)
				(*inNode).right.parent = inNode
				return nil
			}else{
				inNode = (*inNode).right
				//continue
			}
		case (*inNode).Key < key:
			if (*inNode).left == nil{
				(*inNode).left = createNode(key)
				(*inNode).left.parent = inNode
				return nil
			}else{
				inNode = (*inNode).left
				//continue
			}
		}

		}
	}

func (t *Tree) Pop(key int) (*Node, error){
	//foliage
	node, err := t.Search(key)
	if  err != nil {
		return nil, fmt.Errorf("\"%d the key exists \", key")
	} else if (*node).left == nil && nil == (*node).right{
		if (*node).parent != nil{
			if (*node).parent.left == node {
				(*node).parent.left = nil
			}else{
				(*node).parent.right = nil
			}
		}else {
			t.Node = nil
		}
	} else if !((*node).left != nil  && (*node).right != nil){
		err = t.deleteNilNotNil(node)
		if err != nil{
			return nil, fmt.Errorf("ALARM")
		}
	} else {
		err = t.deleteNotNilNotNil(node)
		if err != nil{
			return nil, fmt.Errorf("ALARM")
		}
	}
	return node, nil
}

func (t *Tree)minElTree(node *Node) *Node{
	if (*node).left == nil {
		return node
	} else {
		return t.minElTree(((*node).left))
	}
}

func (t *Tree)maxElTree(node *Node) *Node{
	if (*node).right == nil {
		return node
	} else {
		return t.minElTree(((*node).right))
	}
}

func (t *Tree)deleteNotNilNotNil(node *Node) error{
	minEl := t.minElTree(node.right)
	minEl, _ = t.Pop((*minEl).Key)
	(*minEl).parent, (*minEl).left, (*minEl).right = (*node).parent, (*node).left, (*node).right
	return nil
}

func (t *Tree)deleteNilNotNil(node *Node) error{
	if (*node).left != nil {
		if (*node).parent.left == node{
			(*node).parent.left = node.left
		}else{
			(*node).parent.right = node.left
		}
	}else{
		if (*node).parent.left == node{
			(*node).parent.left = node.right
		}else{
			(*node).parent.right = node.right
		}
	}
	return nil
}

func (t *Tree) Search(key int) (*Node, error){
	var recur func (node *Node) (*Node, error)
	recur = func (node *Node) (*Node, error){
		if (*node).Key == key{
			return node, nil
		}
		if (*node).Key > key{
			if (*node).right == nil{
				return nil, fmt.Errorf("Key not found")
			}
			return recur((*node).right)
		} else {
			if (*node).left ==nil {
				return nil, fmt.Errorf("Key not found")
			}
			return recur((*node).left)
		}
	}

	return recur(t.Node)
}

func createNode(key int)*Node{
	return &Node{key, nil,nil, nil}
}

func (t *Tree) PrintTree(){
	var printNode func (node *Node)
	printNode = func(node *Node) {
		if (*node).left == nil && (*node).right == nil{
			fmt.Printf("%d ", (*node).Key)
			return
		}else{
			fmt.Printf("%d ", (*node).Key)
			if (*node).left != nil{
				printNode((*node).left)
			}
			if (*node).right != nil{
				printNode((*node).right)
			}

		}

	}
	printNode(t.Node)
	fmt.Println()
}

