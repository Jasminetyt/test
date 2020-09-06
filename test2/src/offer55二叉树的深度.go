package main

import (
	"container/list"
	"fmt"
	"test2/tree"
)

func main() {
	var (
		tr *tree.Tree
		nums=[]interface{}{1,2,3,4,5,nil,6,nil,nil,7}
		root=new(tree.Tree)
		)
	root=tr.InitTree(nums,root)
	//树的深度
	fmt.Println(tr.TreeDepth(root))
	//判断是否为平衡二叉树
	fmt.Println(tr.IsBalanced1(root))
	depth:=0
	fmt.Println(tr.IsBalanced2(root,&depth))
	printTreeDepth()
}

type offer55Tree struct {
	value interface{}
	left *offer55Tree
	right *offer55Tree
}

func createoffer55Tree(nums []interface{}) *offer55Tree {
	root:=new(offer55Tree)
	root.value=nums[0]
	var ls=list.New()
	ls.PushBack(root)
	var left,right *offer55Tree
	for i:=1;i<len(nums);{
		left=createoffer55Node(nums[i],ls)
		i++
		if i<len(nums){
			right=createoffer55Node(nums[i],ls)
			i++
		}
		node:=ls.Front()
		node.Value.(*offer55Tree).left=left
		node.Value.(*offer55Tree).right=right
		ls.Remove(node)
	}
	return root
}
func createoffer55Node(value interface{},ls *list.List) *offer55Tree {
	var node *offer55Tree
	if value!=nil{
		node=new(offer55Tree)
		node.value=value
	}else {
		node=nil
	}
	ls.PushBack(node)
	return node
}
func getTreeDepth(root *offer55Tree) int {
	if root==nil{
		return 0
	}
	left:=getTreeDepth(root.left)
	right:=getTreeDepth(root.right)
	if left>right{
		return left+1
	}else {
		return right+1
	}
}
func printTreeDepth()  {
	var nums=[]interface{}{1,2,3,4,5,nil,6,nil,nil,7}
	root:=createoffer55Tree(nums)
	fmt.Println(getTreeDepth(root))
	var depth int
	fmt.Println(isBalance(root,&depth))
}

func isBalance(root *offer55Tree,depth *int) bool{
	if root==nil{
		*depth=0
		return true
	}
	var left,right int
	if isBalance(root.left,&left)&&isBalance(root.right,&right){
		d:=left-right
		if d<=1 || d >=-1{
			if left>right{
				*depth=left+1
			}else {
				*depth=right+1
			}
			return true
		}
	}
	return false
}
