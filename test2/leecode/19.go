package main

import "fmt"

type ListNode struct {
	    Val int
	    Next *ListNode
	 }
func main() {
	var head *ListNode
	head=new(ListNode)
	head.Val=1
	head.Next=nil
	fmt.Println(removeNthFromEnd(head,1))
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head==nil {
		return nil
	}
	var h,p1,p2=head,head,head
	for i:=0;i<n;i++{
		p1=p1.Next
	}
	if p1==nil {//删除的节点为头节点
		h=h.Next
		return h
	}
	for p1.Next!=nil{
		p1=p1.Next
		p2=p2.Next
	}
	temp:=p2.Next
	if temp!=nil {
		temp=temp.Next
	}
	p2.Next=temp
	return h
}

