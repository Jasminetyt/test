package main

import (
	"fmt"
	"test2/list"
	"test2/tree"
	list2"container/list"
)

func main() {
	var nums =[]interface{}{10,6,14,4,8,12,16}
	var tr,root *tree.Tree //<nil>
	var ll *list.LinkedList
	var ls =new(list.LinkedList)
	//root=new(tree.Tree) //将root初始化成树的结构 <<nil> <nil> <nil>>
	root=tr.InitBinaryNodeTree(nums,root)
	tr.TreeConvertList(root,ls)
	ll.TraceList(ls.Next) //将二叉搜索树构建成双向链表第一个节点为空
	fmt.Println("")
	for ls.Next!=nil {
		ls=ls.Next
	}
	for ls.Pre!=nil  {
		fmt.Printf("%v ",ls.Value)
		ls=ls.Pre
	}
	fmt.Println(" ")
	convert2()

}

type offer36Tree struct {
	value interface{}
	left *offer36Tree
	right *offer36Tree
}

func convert2(){
	var num=[]interface{}{10,6,14,4,8,12,16}
	root:=createoffer36Tree(num)
	ls:=convertTreeToList(root)
	for ls.left!=nil{
		ls=ls.left
	}
	for ls!=nil{
		fmt.Printf("%d " ,ls.value.(int))
		ls=ls.right
	}
}
func createoffer36Tree(num []interface{})*offer36Tree{
	var ls=list2.New()
	root:=new(offer36Tree)
	root.value=num[0]
	ls.PushBack(root)
	var left,right *offer36Tree
	for i:=1;i<len(num);{
		if num[i]!=nil{
			left=new(offer36Tree)
			left.value=num[i]
		}else {
				left=nil
		}
		ls.PushBack(left)
		i++
		if i<len(num)&&num[i]!=nil{
			right=new(offer36Tree)
			right.value=num[i]
		}else {
			right=nil
		}
		ls.PushBack(right)
		i++
		node:=ls.Front()
		node.Value.(*offer36Tree).left=left
		node.Value.(*offer36Tree).right=right
		ls.Remove(node)
		}
	return root
}
func convertTreeToList(root *offer36Tree) *offer36Tree {
	if root==nil{
		return nil
	}
	left:=convertTreeToList(root.left)
	for left!=nil&& left.right!=nil{
		left=left.right
	}
	root.left=left
	if left!=nil{
		left.right=root
	}
	right:=convertTreeToList(root.right)
	for right!=nil&&right.left!=nil{
		right=right.left
	}
	root.right=right
	if right!=nil{
		right.left=root
	}
	return root

}

