package Graf

import (
	"fmt"
)

type Graf struct {
	tops map[string]*Node
}

func (graf *Graf) BfsPrint(key string){
	used := make(map[*Node]interface{})
	que := NewQue()
	que.Push(graf.tops[key])
	fmt.Println("Begin")
	for nd, err := que.Pop(); err == nil; nd, err = que.Pop(){
		if _, ok := used[nd.(*Node)]; ok {
			continue
		}
		used[nd.(*Node)] = nil
		for _, node := range nd.(*Node).tops{
			que.Push(node)
		}
		fmt.Println(nd.(*Node).key)
	}

}

func (graf *Graf) SearchBfs(KeySrc, KeyDst string) ([]string, error){
	used := make(map[*Node]interface{})
	edge := make(map[string][]string)
	que := NewQue()
	ndSrc, ok := graf.tops[KeySrc]
	if !ok{
		return nil, fmt.Errorf("Node %s not Faund", KeySrc)
	}
	if _, ok := graf.tops[KeyDst]; !ok{
		return nil, fmt.Errorf("Node %s not Faund", KeyDst)
	}
	que.Push(ndSrc)
	edge[ndSrc.key] = append(edge[ndSrc.key], ndSrc.key)
	for nd, err := que.Pop(); err == nil; nd, err = que.Pop(){
		if nd.(*Node).key == KeyDst{
			return edge[nd.(*Node).key], nil
		}
		for _, node := range nd.(*Node).tops{
			if _, ok := used[node]; ok {
				continue
			}
			que.Push(node)
			edge[node.key] = append(edge[nd.(*Node).key], node.key)
		}
		used[nd.(*Node)] = nil
	}
	return nil, fmt.Errorf("Node not exists")
}

func (graf *Graf)AddTopToTop(key1, key2 string)error{
	_, ok1 := graf.tops[key1]
	if !ok1 {
		graf.tops[key1] = NewNode(key1)
	}
	_, ok2 := graf.tops[key2]
	if !ok2 {
		graf.tops[key2] = NewNode(key2)
	}
	graf.tops[key1].tops = append(graf.tops[key1].tops, graf.tops[key2])
	graf.tops[key2].tops = append(graf.tops[key2].tops, graf.tops[key1]) //Без этого будет ориентриованый граф
	return nil
}

func NewGraf()*Graf{
	return &Graf{make(map[string]*Node)}
}


