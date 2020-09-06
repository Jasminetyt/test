package main

import (
	"test2/queue"
	"test2/tree"
)

func main() {
	var (
		tr *tree.Tree
		tre=new(tree.Tree)
		qu=new(queue.Queue)
		nums=[]interface{}{3,4,2,1,6,4}
		)
	tre=tr.InitTree(nums,tre)
	tr.PrintTreeByLine(tre)
	q:= &qu
	//前序遍历序列化二叉树
	tr.Serialize(tre,q)
	//前序遍历反序列化二叉树
	root2:=tr.DeSerialize(qu)
	tr.PrintTreeByLine(root2)
}
