package point

type ListNode struct {
	Value int
	Next *ListNode
}

//初始化指针
func InitPoint(head *ListNode,num []int)  (*ListNode , *ListNode) {
	var listnode =head
	var temp *ListNode
	//有头节点的尾插法
	//for i := 0;i<len(num);i++{
	//	temp =new(ListNode)
	//	temp.value=num[i]
	//	fmt.Println(temp.value)
	//	temp.next=nil
	//	head.next=temp
	//	head=temp
	//}

	//没有头节点的尾插法
	head.Value=num[0]
	head.Next=nil
	for i := 1;i<len(num) ;i++  {
		temp=new(ListNode)
		temp.Value=num[i]
		temp.Next=nil
		head.Next=temp
		head=temp
	}

	return listnode ,head //返回头节点和尾节点
}
//尾插法
func Intihailpoint(head **ListNode,num []int)  {
	if len(num)==0 {
		return
	}
	(*head).Value=num[0]
	(*head).Next=nil
	node := *head
	for i:=1;i<len(num) ;i++  {
		temp := new(ListNode)
		temp.Value=num[i]
		temp.Next=nil
		node.Next=temp
		node=node.Next
	}

}
//头插法
func Initheadpoint(head **ListNode,num []int)  {
	if len(num)==0 {
		return
	}
	var length=len(num)-1
	(*head).Value=num[length]
	(*head).Next=nil
	for i := length-1;i>=0;i-- {
		temp := new(ListNode)
		temp.Value=num[i];
		temp.Next=*head
		*head=temp
	}
}

func FindKthToTail(head *ListNode,k int) *ListNode {
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

