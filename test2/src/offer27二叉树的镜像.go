package main

import (
	list2 "container/list"
	"fmt"
	"test2/tree"
)

func main() {
	var num=[]interface{}{2,nil,3,nil,nil,nil,4}
	var root = new(tree.Tree)
	var tr tree.Tree
	tr.InitTree(num,root)
	//tr.LevelOrderTrace(root)
	tr.MirrorRecursively(root)
	//fmt.Println("------------------------")
	//tr.LevelOrderTrace(root)
	mirrorRecursively()
}
type offer27Tree struct {
	value interface{}
	left *offer27Tree
	right *offer27Tree
}

func mirrorRecursively(){
	var num=[]interface{}{8,8,7,9,2,nil,nil,nil,nil,4,7}
	root:=offer27TreeCreate(num)
	transferMirrorTree(root)
	traceTree(root)
}

func offer27TreeCreate(num []interface{})*offer27Tree{
	var list=list2.New()
	var root=new(offer27Tree)
	root.value=num[0]
	list.PushBack(root)
	var left,right *offer27Tree
	for i:=1;i<len(num);{
		left=new(offer27Tree)
		left.value=num[i]
		list.PushBack(left)
		i++
		if i<len(num){
			right=new(offer27Tree)
			right.value=num[i]
			list.PushBack(right)
			i++
		}else {
			right=nil
		}
		node:=list.Front()
		node.Value.(*offer27Tree).left=left
		node.Value.(*offer27Tree).right=right
		list.Remove(node)
	}
	return root
}

func transferMirrorTree(root *offer27Tree){
	if root!=nil{
		temp:=root.left
		root.left=root.right
		root.right=temp
		transferMirrorTree(root.left)
		transferMirrorTree(root.right)
	}
}

func traceTree(root *offer27Tree){
	if root!=nil{
		fmt.Println(root.value)
		traceTree(root.left)
		traceTree(root.right)
	}

}
