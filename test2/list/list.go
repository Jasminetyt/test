package list

type LinkedList struct {
	Value interface{}
	Pre *LinkedList
	Next *LinkedList
}

func (list *LinkedList) InitNode(num interface{}) *LinkedList {
	node := new(LinkedList)
	node.Value=num
	node.Next=nil
	node.Pre=nil
	return node
}
//获取链表长度
func (list *LinkedList) Length(head *LinkedList) int {
	count :=0
	if head==nil{
		return count
	}
	for head !=nil{
		head=head.Next
		count++
	}
	return count
}
