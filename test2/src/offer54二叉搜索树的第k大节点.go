package main

import (
	"container/list"
	"fmt"
	"test2/tree"
)

func main() {
	var tr *tree.Tree
	var node=new(tree.Tree)
	var nums=[]interface{}{5,3,7,2,4,6,8}
	node=tr.InitBinaryNodeTree(nums,node)
	k:=3
	if k==0 || node==nil{
		return
	}
	target:=kthNode(node,&k)
	fmt.Println(tr.PrintNodeValue(target))
}
func kthNode(root *tree.Tree ,k *int) (node *tree.Tree) {
	node=nil
	if flag,left:=root.LeftNodeIsExist(root);flag{
		node=kthNode(left,k)
	}
	if node==nil{
		if *k==1{
			node=root
		}
		*k--
		if flag,right := root.RightNodeIsExist(root);flag {
			node=kthNode(right,k)
		}
	}
	return node
}
type offer54tree struct {
	value interface{}
	left *offer54tree
	right *offer54tree
}
func createOffer54Tree(nums []interface{}) *offer54tree {
	var root=new(offer54tree)
	root.value=nums[0]
	var ls=list.New()
	ls.PushBack(root)
	var left,right *offer54tree
	for i:=1;i<len(nums);{
		if nums[i]!=nil{
			left=new(offer54tree)
			left.value=nums[i]
		}else {
			left=nil
		}
		ls.PushBack(left)
		i++
		if i<len(nums)&&nums[i]!=nil{
			right=new(offer54tree)
			right.value=nums[i]
		}else {
			right=nil
		}
		ls.PushBack(right)
		i++
		node:=ls.Front()
		node.Value.(*offer54tree).left=left
		node.Value.(*offer54tree).right=right
		ls.Remove(node)
	}
	return root
}
func kthNode2(root *offer54tree,k *int)*offer54tree{
	var target *offer54tree
	target=nil
	if root.left!=nil{
		target=kthNode2(root.left,k)
	}
	if target==nil{
		if *k==1{
			target=root
		}
		*k--
	}
	if target==nil&&root.right!=nil{
		target=kthNode2(root.right,k)
	}
	return target
}

