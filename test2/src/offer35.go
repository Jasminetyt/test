package main

import "fmt"

type complexListNode struct {
	value int
	next *complexListNode
	sibling  *complexListNode
}
func main() {
	var nums =[]int{1,2,3,4,5,6,7}
	list := initcomlexlist(nums)
	node:=list
	for node != nil{
		fmt.Printf("list :value=%d , siblingvalue=%d \n",node.value,node.sibling.value)
		node=node.next
	}
	fmt.Println("-----------------------------")
	list=cloneNodes(list)
	list=connectSiblingNodes(list)
	list1,list2:= reConnectNodes(list)
	for list1 != nil{
		fmt.Printf("list1 :value=%d , siblingvalue=%d  \n",list1.value,list1.sibling.value)
		list1=list1.next
	}
	fmt.Println("-----------------------------")
	for list2 != nil{
		fmt.Printf("list2 :value=%d , siblingvalue=%d  \n",list2.value,list2.sibling.value)
		list2=list2.next
	}
}
func initcomlexlist(nums []int) *complexListNode {
	list :=initlistnode(nums[0])
	node := list
	for i := 1; i < len(nums);i++ {
		node.next=initlistnode(nums[i])
		node.next.sibling=node
		node=node.next
	}
	list.sibling=node
	return list
}
func initlistnode(num int) *complexListNode {
	temp := new(complexListNode)
	temp.value=num
	temp.next=nil
	temp.sibling=nil
	return temp
}
//1.往复杂链表中加入节点
func cloneNodes(list *complexListNode) *complexListNode {
	node := list
	for node!=nil{
		temp:=initlistnode(node.value)
		temp.next=node.next
		node.next=temp
		node=temp.next
	}
	return list
}
//2.往复制后的链表中添加sibling指向
func connectSiblingNodes(list *complexListNode) *complexListNode {
	node := list
	for node !=nil {
		temp :=node.sibling
		node.next.sibling=temp.next
		node=node.next.next
	}
	return list
}
//3.将复制完成的链表分裂成两条链表
func reConnectNodes(list *complexListNode) (*complexListNode , *complexListNode){
	var (
		list1 ,list2 =list,list.next
		count=0
	)
	list=list2.next
	node1,node2 := list1,list2
	for list != nil {
		if count%2==0 {
			node1.next=list
			count++
			list=list.next
			node1=node1.next
		}else {
			node2.next=list
			count++
			list=list.next
			node2=node2.next
		}
	}
	node1.next=nil
	return list1,list2
}
