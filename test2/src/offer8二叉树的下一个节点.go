package main

import (
	"container/list"
	"fmt"
)

type tttt struct {
	value string
	parent *tttt
	left *tttt
	right *tttt
}
func main() {
	root := initttttnode("a",nil)
	var nums=[]string{"b","c","d","e","f","g","","","h","i"}
	var ls list.List
	ls.PushBack(root)

	for i:=0;i<len(nums);{
		parent := ls.Remove(ls.Front()).(*tttt)
		temp1:=initttttnode(nums[i],parent)
		i++
		var temp2 *tttt
		if i<len(nums){
			temp2=initttttnode(nums[i],parent)
			i++
		}else {
			temp2=nil
		}
		parent.left=temp1
		parent.right=temp2
		ls.PushBack(temp1)
		ls.PushBack(temp2)
	}
	fmt.Println(getNext(root).value)

}
func initttttnode(value string ,parent *tttt)*tttt{
	if value==""{
		return nil
	}
	node := new(tttt)
	node.value=value
	node.right=nil
	node.left=nil
	node.parent=parent
	return node
}
func getNext(node *tttt) *tttt {
	if node==nil {
		return node
	}
	if node.right!=nil{
		temp:=node.right
		for temp.left!=nil{
			temp=temp.left
		}
		return temp
	}else if node.parent!=nil {
		temp:=node
		p:=node.parent
		for p!=nil&&temp==p.right{
			temp=p
			p=p.parent
		}
		return p
	}
	return nil
}

