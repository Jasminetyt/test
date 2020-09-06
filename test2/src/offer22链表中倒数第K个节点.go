package main

import (
	list2 "container/list"
	"fmt"
	"test2/point"
)

func main() {
	var num =[]int{1,2,3,4,5,6}
    var head=new(point.ListNode)
    var head2 **point.ListNode
    head2=&head
    point.Intihailpoint(head2,num)
    result := findKthToTail(head,2)
    fmt.Println(result.Value)
    var list =list2.New()
    for i:=0;i<len(num);i++{
    	list.PushBack(num[i])
	}
    fmt.Println(findKthToTail2(list,2).Value)
}
func findKthToTail(head *point.ListNode,k int) *point.ListNode {
	if head == nil || k == 0 {
		return nil
	}
	head2 := head
	for i:=0;i<k-1;i++ {
		if head2.Next != nil{ //如果为head2!=nil 则head会走过
			head2=head2.Next
		}else {
			return nil
		}
	}
	for head2.Next!=nil { //如果为head2!=nil 则head会走过
		head2=head2.Next
		head=head.Next
	}
	return head
}
func findKthToTail2(list *list2.List,k int) *list2.Element {
	node1,node2:=list.Front(),list.Front()
	for ;k>0&&node1!=nil;k--{
		node1=node1.Next()
	}
	if k>0&&node1==nil{
		return nil
	}
	for node1!=nil{
		node1=node1.Next()
		node2=node2.Next()
	}
	return node2
}
