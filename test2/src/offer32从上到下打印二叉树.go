package main

import (
	"bufio"
	list2 "container/list"
	"fmt"
	"os"
	"test2/tree"
)
//接口由type和data两部分组成，只有两者均为nil时，该接口才为nil。
//var a interface{} = (*int)nil 此时a!=nil
func main() {
	var str []interface{}
	var tr =new(tree.Tree)
	scan := bufio.NewScanner(os.Stdin)
	scan.Split(bufio.ScanWords)
	fmt.Println("请输入")
	for scan.Scan()  {
		if scan.Text()=="end" {
			break
		}
		if scan.Text()=="nil" {
			str=append(str,nil)
		}else {
			str=append(str,scan.Text())
		}
	}
	//for i,_:= range str{
	//	fmt.Println(reflect.TypeOf(interface{}(str[i])))
	//vi :=reflect.ValueOf(str[i])//判断接口的动态值类型
	//if vi.Kind()==reflect.Ptr{ //如果vi不是指针类型，则会报错
	//	fmt.Println(vi.IsNil())//若vi为指针同时值为nil返回true，否者返回false
	//}else {
	//	fmt.Println(false)
	//}
	//}
	root:=tr.InitTree(str,tr)
	tr.LevelOrderTrace(root)
	fmt.Print("\n") 
	tr.PrintTreeByLine(root) //题目二
	tr.PrintTreeByZhi(root)//题目三
	printFromTopToBottom()
	fmt.Println()
	printFromTopToBottom2()
	fmt.Println()
	print()
}

type offer32Tree struct {
	value interface{}
	left *offer32Tree
	right *offer32Tree
}
//不分行从上到下打印二叉树
func printFromTopToBottom(){
	var num=[]interface{}{8,8,7,9,2,nil,nil,nil,nil,4,7}
	root:=createTree(num)
	var list=list2.New()
	list.PushBack(root)
	for list.Len()>0{
		node:=list.Front()
		fmt.Printf("%d ",node.Value.(*offer32Tree).value.(int))
		if node.Value.(*offer32Tree).left!=nil{
			list.PushBack(node.Value.(*offer32Tree).left)
		}
		if node.Value.(*offer32Tree).right!=nil{
			list.PushBack(node.Value.(*offer32Tree).right)
		}
		list.Remove(node)
	}
}
func createTree(num []interface{})*offer32Tree{
	var list=list2.New()
	root:=new(offer32Tree)
	root.value=num[0]
	list.PushBack(root)
	var left,right *offer32Tree
	for i:=1;i<len(num);{
		if num[i]!=nil{
			left=new(offer32Tree)
			left.value=num[i]
			list.PushBack(left)
		}else {
			left=nil
		}
		i++
		if i<len(num)&&num[i]!=nil{
			right=new(offer32Tree)
			right.value=num[i]
			list.PushBack(right)

		}else {
				right=nil
		}
		i++
		node:=list.Front()
		node.Value.(*offer32Tree).left=left
		node.Value.(*offer32Tree).right=right
		list.Remove(node)
	}
	return root
}

//分行从上到下打印二叉树
func printFromTopToBottom2()  {
	var num=[]interface{}{8,8,7,9,2,nil,nil,nil,nil,4,7}
	root:=createTree(num)
	var tobePrint,nexlevel=1,0
	var list=list2.New()
	list.PushBack(root)
	for list.Len()>0{
		node:=list.Front()
		fmt.Printf("%d ",node.Value.(*offer32Tree).value)
		list.Remove(node)
		tobePrint--
		if node.Value.(*offer32Tree).left!=nil{
			list.PushBack( node.Value.(*offer32Tree).left)
			nexlevel++
		}
		if node.Value.(*offer32Tree).right!=nil{
			list.PushBack(node.Value.(*offer32Tree).right)
			nexlevel++
		}
		if tobePrint==0{
			fmt.Println(" ")
			tobePrint=nexlevel
			nexlevel=0
		}
	}

}

//之字形打印二叉树
func print()  {
	var num=[]interface{}{8,8,7,9,2,nil,nil,nil,nil,4,7}
	root:=createTree(num)
	var ls1=list2.New()
	var ls2=list2.New()
	ls1.PushBack(root)
	var flag=1
	var tobeprint,nextlevel=1,0
	for ls1.Len()>0||ls2.Len()>0{
		if flag==1{
			node:=ls1.Back()
			fmt.Printf("%d ",node.Value.(*offer32Tree).value)
			tobeprint--
			ls1.Remove(node)
			if node.Value.(*offer32Tree).left!=nil{
				ls2.PushBack(node.Value.(*offer32Tree).left)
				nextlevel++
			}
			if node.Value.(*offer32Tree).right!=nil{
				ls2.PushBack(node.Value.(*offer32Tree).right)
				nextlevel++
			}
		}else {
			node:=ls2.Back()
			fmt.Printf("%d ",node.Value.(*offer32Tree).value)
			tobeprint--
			ls2.Remove(node)
			if node.Value.(*offer32Tree).right!=nil{
				ls1.PushBack(node.Value.(*offer32Tree).right)
				nextlevel++
			}
			if node.Value.(*offer32Tree).left!=nil{
				ls1.PushBack(node.Value.(*offer32Tree).left)
				nextlevel++
			}
		}
		if tobeprint==0{
			tobeprint=nextlevel
			nextlevel=0
			flag=-flag
			fmt.Println(" ")
		}
	}
}

