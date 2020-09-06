package main

import (
	"container/list"
	"fmt"
	list2 "test2/list"
	"test2/point"
)

func main() {
	var head2 **point.ListNode
	var head=new(point.ListNode)
	var num=[]int{1,1,2,2,3,4,4,5,5,6,7,7,8,8}
	head,_=point.InitPoint(head,num)
	head2=&head
	deleteDuplication(head2)
	head=*head2
	for head !=nil {
		fmt.Printf("value=%d ",head.Value)
		head=head.Next
	}
	fmt.Println("")
	var ls=list.New()
	for i:=0;i<len(num);i++{
		ls.PushBack(num[i])
	}
	deleteDuplication2(ls)
	deleteDuplication3()
}
//删除重复节点
func deleteDuplication(head2 **point.ListNode)  {
	var(
		node = *head2
		preNode ,nextNode,temp *point.ListNode
		flag=0 //标记是否需要设置为头节点
		value int
	)
	for node != nil  {
		value=node.Value
		preNode=node
		node=node.Next
		for node!=nil&&node.Value==value  {
			preNode=node
			node=node.Next
		}
		if node==nil {
			return
		}
		nextNode=node.Next
		if nextNode ==nil{
			preNode.Next=nil
			temp.Next=node
			temp=temp.Next
			temp.Next=nil
		}
		if nextNode!=nil&&nextNode.Value!=node.Value{
			flag++
			preNode.Next=nil
			preNode=node
			node=node.Next
			if flag>1 {
				temp.Next=preNode
				temp=temp.Next
				temp.Next=nil
			}
		}
		if flag==1 { //当为1时则设为头节点
			*head2=preNode
			temp=*head2
			temp.Next=nil
			flag++
		}

	}

}

func deleteDuplication2(ls *list.List){
	var nextNode *list.Element
	for node:=ls.Front();node!=nil;{
		nextNode=node.Next()
		if nextNode!=nil&& nextNode.Value.(int)==node.Value.(int){
			for nextNode.Value.(int)==node.Value.(int){
				ls.Remove(nextNode)
				nextNode=ls.Front().Next()
			}
			ls.Remove(node)
		}
		node=nextNode
	}
	for ll:=ls.Front();ll!=nil;ll=ll.Next(){
		fmt.Println(ll.Value.(int))
	}
}

func deleteDuplication3(){
	var ls *list2.LinkedList
	var num=[]interface{}{1,1,2,2,3,4,4,5,5,6,7,7,8,8}
	head:=ls.InitTailInsertListNoHead(num)
	var preNode,nextNode *list2.LinkedList
	node := head
	for node!=nil{
		nextNode=node.Next
		if nextNode!=nil&& node.Value.(int)==nextNode.Value.(int) {
			for nextNode!=nil&&nextNode.Value.(int)==node.Value.(int){
				node=nextNode
				nextNode=node.Next
			}
		}else {
			if preNode==nil{
				preNode=node
				head=preNode
			}else {
				preNode.Next=node
				preNode=node
			}
		}
		node=nextNode
	}
	preNode.Next=node
	for head!=nil{
		fmt.Println(head.Value.(int))
		head=head.Next
	}

}