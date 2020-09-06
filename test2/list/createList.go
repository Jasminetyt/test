package list

//尾插法初始化不带头节点的链表
func (list *LinkedList) InitTailInsertListNoHead(nums []interface{}) *LinkedList {
	if len(nums)<=0 {
		return nil
	}
	List := list.InitNode(nums[0])
	node := List
	for i:=1;i<len(nums);i++ {
		temp:=list.InitNode(nums[i])
		node.Next=temp
		node=node.Next
	}
	return List
}
//尾插法初始化带头节点的链表
func (list *LinkedList) InitTailInsertList(nums []interface{}) *LinkedList {
	if len(nums)<=0 {
		return nil
	}
	List := new(LinkedList)
	node := List
	for i:=0; i<len(nums);i++  {
		temp := list.InitNode(nums[i])
		node.Next=temp
		node=node.Next
	}
	return List
}
//头插法初始化不带头节点的链表
func (list *LinkedList) InitHeadInsertListNoHead(nums []interface{}) *LinkedList  {
	if len(nums)<=0 {
		return nil
	}
	List := list.InitNode(nums[0])
	for i:=1;i<len(nums);i++  {
		temp := list.InitNode(nums[i])
		temp.Next=List
		List=temp
	}
	return List
}
//头插法初始化带头节点的链表
func (list *LinkedList) InitHeadInsert(nums []interface{}) *LinkedList  {
	if len(nums)<=0 {
		return nil
	}
	List := new(LinkedList)
	for i:=0;i<len(nums) ; i++ {
		temp := list.InitNode(nums[i])
		temp.Next=List.Next
		List.Next=temp
	}
	return List
}
//创建带有公共节点的链表
func (list *LinkedList) InitLinkListWithCommonNode(nums1,nums2,nums3 []interface{}) (list1,list2 *LinkedList) {
	list1=list.InitTailInsertListNoHead(nums1)
	list2=list.InitTailInsertListNoHead(nums2)
	list3 := list.InitTailInsertListNoHead(nums3)
	temp1,temp2 := list1,list2

	for temp1!=nil&&temp1.Next!=nil{
		temp1=temp1.Next
	}
	for temp2!=nil&&temp2.Next!=nil{
		temp2=temp2.Next
	}
	if temp1==nil{
		temp1=list3
	}else {
		temp1.Next=list3
	}
	if temp2==nil{
		temp2=list3
	}else {
		temp2.Next=list3
	}
	return list1,list2
}
