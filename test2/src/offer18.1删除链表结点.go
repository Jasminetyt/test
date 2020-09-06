package main

import (
	"fmt"
	"test2/point"
)
func main() {
	var head=new(point.ListNode)
	var tail *point.ListNode
	var num=[]int{1,1,2,2,3,4,4,5,5,6,7,7,8}
	head,tail=point.InitPoint(head,num)
	//var dnode=head.next
	deleteNode(head,tail) //删除尾节点
	for head !=nil {
		fmt.Printf("value=%d ",head.Value)
		head=head.Next
	}
}


//若删除节点为尾节点，则需从头到位遍历o(n)
//若删除节点为中间节点，只需把该节点指向的下一个值复制到当前节点，删除下个节点即可o(1)
//若删除节点为头节点，将头节点置为nil
//[（n-1）*o(1)+o(n) ] / n = o(1)
func deleteNode(head *point.ListNode,node *point.ListNode)  {
	if head==node && head.Next==nil{
		head=nil
	}else if node.Next!=nil {
		var nextnode=node.Next
		node.Value=nextnode.Value
		node.Next=nextnode.Next
	}else {
		var temp *point.ListNode
		for head!=node  {
			temp=head
			head=head.Next
		}
		temp.Next=nil
	}
}

