package main

import (
	"fmt"
	"test2/list"
	"test2/point"
)

func main() {
	var num = []int{1,2,3,4,5,6}
	var head=new(point.ListNode)
	var head2 = &head
	point.Intihailpoint(head2,num)
	result := reverseList(head)
	for result!=nil{
		fmt.Println(result.Value)
		result=result.Next
	}
	reverseList2()
}
func reverseList(head *point.ListNode) *point.ListNode {
	if head==nil {
		return nil
	}
	var(
		prenode *point.ListNode
		node *point.ListNode
		nextnode *point.ListNode
	)
	prenode=nil
	node=head
	nextnode=node.Next
	for nextnode!=nil {
		node.Next=prenode
		prenode=node
		node=nextnode
		nextnode=node.Next
	}
	node.Next=prenode
	return node
}

func reverseList2(){
	var nums=[]int{1,2,3,4}
	var head =new(list.LinkedList)
	if len(nums)==0{
		return
	}
	head.Value=nums[0]
	node:=head
	for i:=1;i<len(nums);i++{
		temp:=new(list.LinkedList)
		temp.Value=nums[i]
		node.Next=temp
		node=temp
	}
	var p1,p2,p3 *list.LinkedList
	p1=nil
	p2=head
	for p2!=nil{
		p3=p2.Next
		p2.Next=p1
		p1=p2
		p2=p3
	}
	fmt.Println("----------------")
	for p1!=nil{
		fmt.Println(p1.Value)
		p1=p1.Next
	}
}