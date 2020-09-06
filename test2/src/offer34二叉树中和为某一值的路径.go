package main

import (
	"bufio"
	list2 "container/list"
	"fmt"
	"os"
	"strconv"
	"test2/tree"
)

func main() {
	var n int
	fmt.Println("输入和")
	_,_=fmt.Scan(&n)
	fmt.Println("输入树的节点值")
	var nums []interface{}
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		if scan.Text()=="end" {
			break
		}
		if scan.Text()=="nil"{
			nums=append(nums,nil)
		}else {
			num,_:= strconv.Atoi(scan.Text())
			nums=append(nums,num)
		}
	}
	var root *tree.Tree
	if len(nums)==0 {
		root=nil
	}else {
		root =new (tree.Tree)
		root=root.InitTree(nums,root)
	}
	root.FindPath(root,n,0)
	fmt.Println("----------------")
	findpath2()
}
type offer34Tree struct {
	value interface{}
	left *offer34Tree
	right *offer34Tree
}
func findpath2()  {
	root:=offer34createTree()
	var sum=22
	var num []interface{}
	find2(root,sum,num,0)

}
func offer34createTree() *offer34Tree {
	var num=[]interface{}{10,5,12,4,7,nil,13}
	var list=list2.New()
	root:=new(offer34Tree)
	root.value=num[0]
	list.PushBack(root)
	var left,right *offer34Tree
	for i:=1;i<len(num);{
		if num[i]!=nil{
			left=new(offer34Tree)
			left.value=num[i]
			list.PushBack(left)
		}else {
			left=nil
		}
		i++
		if i<len(num)&&num[i]!=nil{
			right=new(offer34Tree)
			right.value=num[i]
			list.PushBack(right)
		}else {
			right=nil
		}
		i++
		node:=list.Front()
		node.Value.(*offer34Tree).left=left
		node.Value.(*offer34Tree).right=right
		list.Remove(node)
	}
	return root
}

func find2(root *offer34Tree,sum int,num []interface{},temp int)  {
	if root==nil{
		num=num[:len(num)-1]
		return
	}
	if temp+root.value.(int)==sum{
		num=append(num,root.value)
		fmt.Println(num)
		num=num[:len(num)-1]
		return
	}
	if temp+root.value.(int)>sum{
		num=num[:len(num)-1]
		return
	}
	temp=temp+root.value.(int)
	num=append(num,root.value)
	find2(root.left,sum,num,temp)
	find2(root.right,sum,num,temp)

}
