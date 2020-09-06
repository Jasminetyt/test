package main

import (
	"container/list"
	"fmt"
	stack3 "test2/stack"
)

func main() {
	var stack1 stack3.Stack
	stack:=stack1.New()
	num := []int{3,4,5,2,1,5}
	for i:=0;i<len(num) ;i++  {
		stack.StackWithMinPush(num[i])
	}
	fmt.Println(stack.StackWithMin())
	for i:=0;i<3 ;i++  {
		stack.StackWithMinPop()
	}
	fmt.Println(stack.StackWithMin())

	var os =new(offer30Stack)
	os.init()
	for i:=0;i<len(num);i++{
		os.push(num[i])
	}
	fmt.Println(os.min())
	for i:=0;i<4;i++{
		os.pop()
	}
	fmt.Println(os.min())
}

type offer30Stack struct {
	list1 *list.List
	list2 *list.List
}

func (os *offer30Stack) init(){
	os.list1=list.New()
	os.list2=list.New()
}

func (os *offer30Stack) push(num int){
	os.list1.PushBack(num)
	if os.list2.Len()>0{
		node:=os.list2.Back()
		if node.Value.(int)>num {
			os.list2.PushBack(num)
		}else {
			os.list2.PushBack(node.Value.(int))
		}
	}else {
		os.list2.PushBack(num)
	}
}

func (os *offer30Stack) pop()  {
	var node *list.Element
	if os.list1.Len()>0 {
		node=os.list1.Back()
		os.list1.Remove(node)
		node=os.list2.Back()
		os.list2.Remove(node)
	}
}

func (os *offer30Stack) min()int{
	return os.list2.Back().Value.(int)
}