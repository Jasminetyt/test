package main

import (
	list2 "container/list"
	"fmt"
	"test2/queue"
	"test2/tree"
)
var (
	tr tree.Tree
	que queue.Queue
)
func main() {
	var (
		num1=[]interface{}{2,2,nil,3,nil,nil,nil,4}
		num2=[]interface{}{3,4,nil}
		root1=new(tree.Tree)
		root2=new(tree.Tree)
	)
	tr.InitTree(num1,root1)
	tr.InitTree(num2,root2)
	fmt.Println(doesTree1HasTree2(root1,root2))
	doesTree1HasTree22()
}
func doesTree1HasTree2(root1 ,root2 *tree.Tree) bool {
	var flag=false
	que=tr.Getsamenode(root1,root2)
	defer que.Clear()
	for ! que.IsEmpty() {
		node := interface{}(que.DeQueue()).(*tree.Tree)
		flag=tr.DoesSameStruct(node,root2)
		if flag==true {
			break
		}
	}
	return flag
}
type offer26Tree struct {
	left *offer26Tree
	right *offer26Tree
	value interface{}
}
func doesTree1HasTree22(){
	var num1=[]interface{}{8,8,7,9,2,nil,nil,nil,nil,4,7}
	var num2=[]interface{}{8,9,2}
	root1:=offer26createtree(num1)
	root2:=offer26createtree(num2)

	var list=list2.New()
	samenode(root1,root2,list)
	for node:=list.Front();node!=nil;node=node.Next(){
		if issame(node.Value.(*offer26Tree),root2){
			fmt.Println("true")
			return
		}
	}
	fmt.Println("false")

}
func offer26createtree(num []interface{}) *offer26Tree {
	var list=list2.New()
	root:=new(offer26Tree)
	root.value=num[0]
	list.PushBack(root)

	for i:=1;i<len(num);{
		var right,left *offer26Tree
		left=new(offer26Tree)
		left.value=num[i]
		list.PushBack(left)
		i++
		if i<len(num){
			right=new(offer26Tree)
			right.value=num[i]
			list.PushBack(right)
			i++
		}else {
			right=nil
		}
		node:=list.Front()
		node.Value.(*offer26Tree).left=left
		node.Value.(*offer26Tree).right=right
		list.Remove(node)

	}
	return root
}
func samenode(root1,root2 *offer26Tree,list *list2.List){

	if root1==nil{
		return
	}
	if root1.value!=nil&&root1.value.(int)==root2.value.(int){
		list.PushBack(root1)
	}
	samenode(root1.left,root2,list)
	samenode(root1.right,root2,list)
}
func issame(root1,root2 *offer26Tree) bool {
	if root2==nil{
		return true
	}
	if root2!=nil&&root1==nil{
		return false
	}
	if root1.value==root2.value{
		return issame(root1.left,root2.left)&&issame(root1.right,root2.right)
	}
	return false
}
