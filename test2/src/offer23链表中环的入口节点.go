package main

import (
	list2 "container/list"
	"fmt"
	"test2/point"
)

func main() {
	//构造带环链表
	var(
		num=[]int{1,2,3,4,5,6,7,8}
		head = new(point.ListNode)
		tail *point.ListNode
	)
	head ,tail = point.InitPoint(head,num)
	node := point.FindKthToTail(head,3)
	tail.Next=node
	meetingnode := meetingNode(head)
	if meetingnode==nil {
		fmt.Println("该链表没有环")
	}else {
		fmt.Println("方法1")
		length := countlength(meetingnode)
		result := entryNodeOfLoop(head,length)
		fmt.Println(result)
		fmt.Println("方法2")
		result2 := entryNodeOfLoop2(head,meetingnode)
		fmt.Println(result2)
	}


}
//求该链表是否有环，若有则求出相遇点
func meetingNode(head *point.ListNode) *point.ListNode {
	if head==nil {
		fmt.Println("空链表")
		return nil
	}
	head2 := head
	for head !=nil  {
		head=head.Next
		if head!=nil {
			head=head.Next
		}
		head2=head2.Next
		if head==head2 {
			break
		}
	}
	if head==nil {
		return nil
	}else {
		fmt.Printf("相遇点为：%v\n",head)
		return head
	}
}
//计算环的长度，输入参数为相遇点
func countlength(node *point.ListNode) int {
	var count=1
	temp := node.Next
	for node != temp  {
		temp=temp.Next
		count++
	}
	return count
}
//计算入口点
func entryNodeOfLoop(head *point.ListNode,length int) *point.ListNode  {
	head2 := head
	for length>0 {
		head=head.Next
		length--
	}
	for head!=head2  {
		head=head.Next
		head2=head2.Next
	}
	return  head2
}
//求解入环口的第二种方法
func entryNodeOfLoop2(head *point.ListNode , node *point.ListNode) *point.ListNode {
	for node!=head{
		node=node.Next
		head=head.Next
	}
	return head
}

func meetingNode2(){
	var nums=[]int{1,2,3,4,5,6}
	var list=list2.New()
	for i:=0;i<len(nums);i++{
		list.PushBack(nums[i])
	}
	last:=list.Back()
	first:=list.Front()
	last.Next()=first.Next().Next()//构造带环的链表!!!!!!!无法构造带环的链表，只能构造双向链表

//判断该链表是否有环
	node1,node2:=first,first
	for node2!=nil{
		node1=node1.Next()
		node2=node2.Next().Next()
		if node1==node2{
			break
		}
	}
	if node2!=nil{
		fmt.Println(node1.Value)//找到节点
	}else {
		return
	}

	//寻找入环口
	//解法一
	temp:=first
	for node1!=temp{
		node1=node1.Next()
		temp=temp.Next()
	}
	fmt.Println(temp.Value)
	//解法二
	var count =1
	temp=node2.Next()
	for temp!=node2{
		temp=temp.Next()
		count++ //环长度
	}
	temp1,temp2:=first,first
	for i:=0;i<=count;i++{
		temp1=temp1.Next()
	}
	for temp2!=temp1{
		temp2=temp2.Next()
		temp1=temp1.Next()
	}
	fmt.Println(temp1.Value)
}
