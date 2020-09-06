package main

import (
	"fmt"
	"test2/list"
	list2"container/list"
)
var ls list.LinkedList
func main() {
	var (
		num1=[]interface{}{1,2,3}
	    num2=[]interface{}{1,2,3}
	    num3=[]interface{}{4,5,6 }
	)
	list1,list2 :=ls.InitLinkListWithCommonNode(num1,num2,num3)
	common:=findFirstCommonNode(list1,list2)
	ls.TraceList(common)

	findFirstCommonNode2(num1,num2,num3)
}
func findFirstCommonNode(list1,list2 *list.LinkedList) *list.LinkedList {
	length1 := ls.Length(list1)
	length2 := ls.Length(list2)
	for length1>length2{
		list1=list1.Next
		length1--
	}
	for length2>length1{
		list2=list2.Next
		length2--
	}
	for list1!=list2 {
		list1=list1.Next
		list2=list2.Next
	}
	return list1
}
type offer52Link struct {
	value interface{}
	next *offer52Link
}
func findFirstCommonNode2(nums1,nums2,nums3 []interface{})  {
	ls1,end1:=createLink(nums1)
	ls2,end2:=createLink(nums2)
	ls3,_:=createLink(nums3)
	if end1==nil{
		ls1=ls3
	}else {
		end1.next=ls3
	}
	if end2==nil{
		ls2=ls3
	}else {
		end2.next=ls3
	}
	var l1,l2=list2.New(),list2.New()
	for ls1!=nil{
		l1.PushBack(ls1)
		ls1=ls1.next
	}
	for ls2!=nil{
		l2.PushBack(ls2)
		ls2=ls2.next
	}
	var temp *offer52Link
	for node1,node2:=l1.Back(),l2.Back();node1!=nil&&node2!=nil;node1,node2=node1.Prev(),node2.Prev(){
		if node1.Value.(*offer52Link)!=node2.Value.(*offer52Link){
			temp=node1.Value.(*offer52Link)
			break
		}
	}
	if temp==nil||temp.value==nil{
		fmt.Println("not exist")
		return
	}
	if temp.next!=nil{
		fmt.Println(temp.next.value)
	}else {
		fmt.Println("not exist")
	}
}
func createLink(nums []interface{}) (head,end *offer52Link) {
	if len(nums)==0{
		return nil,nil
	}
	var ls=new(offer52Link)
	ls.value=nums[0]
	temp:=ls
	for i:=1;i<len(nums);i++{
		node:=new(offer52Link)
		node.value=nums[i]
		temp.next=node
		temp=temp.next
	}
	return ls,temp
}
