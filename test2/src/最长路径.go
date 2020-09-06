package main

import (
	"fmt"
	"test2/tree"
)


func main() {
	var nums=[]interface{}{1,2,3,4,5,nil,nil}
	var tr *tree.Tree
	var root=new(tree.Tree)
	root2 :=tr.InitTree(nums,root)
	//tr.LevelOrderTrace(root2)
	//fmt.Println()
	fmt.Println(tr.MaxTreeLength(root2))
}

