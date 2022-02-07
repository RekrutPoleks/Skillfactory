package main

import "Skillfactory/GrafandBTS/BTS"

func main()  {
	//fmt.Println("Hello, playground")
	//graf := Graf.NewGraf()
	//graf.AddTopToTop("A", "B")
	////graf.AddTopToTop("C", "A")
	//graf.AddTopToTop("B", "D")
	//graf.AddTopToTop("C", "D")
	//graf.AddTopToTop("E", "C")
	//graf.AddTopToTop("F", "D")
	//fmt.Println(graf)
	//fmt.Println(graf.SearchBfs("F", "A"))
	tree := BTS.Tree{}
	tree.Push(56)
	tree.Push(43)
	tree.Push(5)
	tree.Push(6)
	tree.Push(123)
	tree.Push(32)
	tree.Push(10)
	tree.PrintTree()
	tree.Pop(43)
	tree.PrintTree()
}
