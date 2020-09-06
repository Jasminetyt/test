package list

import "fmt"

//从前向后遍历链表
func (list *LinkedList) TraceList(head *LinkedList)  {
	if head==nil {
		return
	}
	for head != nil {
		fmt.Printf("%v ",head.Value)
		head=head.Next
	}
	fmt.Println()
}