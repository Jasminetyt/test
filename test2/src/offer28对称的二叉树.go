package main

import (
	list2 "container/list"
	"fmt"
	"test2/tree"
)

func main() {
	var (
		tr tree.Tree
		num=[]interface{}{8,6,6,5,7,7}
		root=new(tree.Tree)
	)
	tr.InitTree(num,root)
	fmt.Println(tr.IsSymmetrical(root))
	symmetricalTree()
}
type offer28Tree struct {
	value interface{}
	left *offer28Tree
	right *offer28Tree
}
func symmetricalTree(){
	var nums=[]interface{}{8,6,6,5,7,7,5}
	root:=offer28CreateTree(nums)
	fmt.Println(isSymmetrical(root,root))
}

func offer28CreateTree(nums []interface{}) *offer28Tree{
	var list=list2.New()
	var root=new(offer28Tree)
	root.value=nums[0]
	list.PushBack(root)
	var left,right *offer28Tree
	for i:=1;i<len(nums);{
		left=new(offer28Tree)
		left.value=nums[i]
		list.PushBack(left)
		i++
		if i<len(nums){
			right=new(offer28Tree)
			right.value=nums[i]
			list.PushBack(right)
			i++
		}else {
				right=nil
		}
		node:=list.Front()
		node.Value.(*offer28Tree).left=left
		node.Value.(*offer28Tree).right=right
		list.Remove(node)
	}
	return root
}

func isSymmetrical(root1,root2 *offer28Tree)bool{
	if root1==nil&&root2==nil{
		return true
	}
	if (root1==nil&&root2!=nil)||(root2==nil&&root1!=nil){
		return false
	}
	if root1.value.(int)==root2.value.(int){
		return isSymmetrical(root1.left,root2.right)&&isSymmetrical(root1.right,root2.left)
	}
	return false
}
