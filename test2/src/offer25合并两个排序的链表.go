package main

import (
	"fmt"
	"test2/list"
	"test2/point"
)

func main() {
var(
	num1 = []int{1,3,6,8,9}
	num2 = []int{2,4,5,7}
)
head1 := new(point.ListNode)
head2 := new(point.ListNode)
point.Intihailpoint(&head1,num1)
point.Intihailpoint(&head2,num2)
head := mergelist(head1,head2)
	for head != nil {
		fmt.Println(head.Value)
		head=head.Next
	}
	mergelist2()
}
func mergelist(head1,head2 *point.ListNode) *point.ListNode {
	if head1 ==nil {
		return head2
	}else if head2 == nil{
		return head1
	}
	var head = new(point.ListNode)
	if head1.Value<=head2.Value {
		head=head1
		head.Next=mergelist(head1.Next,head2)//使用递归
	}else {
		head=head2
		head.Next=mergelist(head1,head2.Next)
	}
	//node := head
	//for head1!=nil && head2 != nil {
	//	if head1.Value<=head2.Value {
	//		node.Next=head1
	//		head1=head1.Next
	//		node=node.Next
	//	}else {
	//		node.Next=head2
	//		head2=head2.Next
	//		node=node.Next
	//	}
	//}
	//for head1!=nil  {
	//	node.Next=head1
	//	head1=head1.Next
	//	node=node.Next
	//}
	//for head2 != nil {
	//	node.Next=head2
	//	head2=head2.Next
	//	node=node.Next
	//}
	 return head
}
func mergelist2(){
	var(
		num1 = []int{1,3,6,8,9}
		num2 = []int{}
	)
	var head3 =new(list.LinkedList)
	if len(num1)==0&&len(num2)==0{
		return
	}
	var head1,head2 *list.LinkedList
	var node1,node2 *list.LinkedList
	if len(num1)!=0{
		head1=new(list.LinkedList)
		node1=head1
		node1.Value=num1[0]
	}
	if len(num2)!=0{
		head2=new(list.LinkedList)
		node2=head2
		node2.Value=num2[0]
	}
	for i:=1;i<len(num1);i++{
		temp:=new(list.LinkedList)
		temp.Value=num1[i]
		node1.Next=temp
		node1=temp
	}
	for i:=1;i<len(num2);i++{
		temp:=new(list.LinkedList)
		temp.Value=num2[i]
		node2.Next=temp
		node2=temp
	}


	node:=head3//头节点为空
	for head1!=nil&&head2!=nil{
		if head1.Value.(int)<head2.Value.(int){
			node.Next=head1
			node=head1
			head1=head1.Next
		}else {
			node.Next=head2
			node=head2
			head2=head2.Next
		}
	}

	for  head1!=nil{
		node.Next=head1
		node=head1
		head1=head1.Next
	}

	for  head2!=nil{
		node.Next=head2
		node=head2
		head2=head2.Next
	}
	head3=head3.Next
	for head3!=nil{
		fmt.Printf("%d ",head3.Value.(int))
		head3=head3.Next
	}
}

