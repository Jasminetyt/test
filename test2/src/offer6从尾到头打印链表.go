package main

import (
	"container/list"
	"fmt"
)

func main() {
	head := list.New()
	stack:=list.New()
	for i:=0;i<6;i++{
		head.PushBack(i)
	}
	for head2:=head.Front();head2!=nil;head2=head2.Next(){
		stack.PushFront(head2.Value)
	}
	for head2 :=stack.Front();head2!=nil;head2=head2.Next(){
		fmt.Println(head2.Value.(int))
	}
}
