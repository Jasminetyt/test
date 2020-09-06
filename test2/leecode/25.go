package main

import (
	list2 "test2/list"
)

func main() {
	var list  list2.LinkedList
	var nums=[]interface{}{1,2,3,4,5}
	head :=list.InitTailInsertListNoHead(nums)
	//list.TraceList(head)
	head=reverseKGroup(head,0)
	list.TraceList(head)
}
func reverseKGroup(head *list2.LinkedList, k int) *list2.LinkedList  {
	if k<=1 {
		return head
	}
	var count int
	var temp,temp1,temp2,p1,p2 *list2.LinkedList
	var flag=true
	for head !=nil{
		temp=head
		count=1
		for head!=nil&&count<k{
			head=head.Next
			count++
		}
		if head!=nil{
			head=head.Next
			temp1,temp2=reverseLinkList(temp,k)
			if flag {
				p1=temp1
				p2=temp2
				flag=false
			}else {
				p2.Next=temp1
				p2=temp2
			}
		}else {
			if flag {
				p1=temp
				flag=false
			}else {
				p2.Next=temp
			}
		}
	}
	return p1
}
func reverseLinkList(p *list2.LinkedList,k int)(head,tail *list2.LinkedList) {
	count:=1
	tail=p
	p2:=p.Next
	var p3 *list2.LinkedList
	for count<k{
		p3=p2.Next
		p2.Next=p
		p=p2
		p2=p3
		count++
	}
	head=p
	tail.Next=nil
	return head,tail
}