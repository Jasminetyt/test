package tree

import (
	"fmt"
	"test2/list"
	"test2/queue"
	"test2/stack"
)
type Tree struct {
		value interface{}
		left *Tree
		right *Tree
	}
var (
	que queue.Queue
	quee = que.New()
	stk stack.Stack
	sta=stk.New()
	ls list.LinkedList
	)


//判断左节点是否存在
func (tree *Tree) LeftNodeIsExist(root *Tree) (bool,*Tree){
	if root.left==nil {
		return false,nil
	}else {
		return true,root.left
	}
}
//判断右节点是否存在
func (tree *Tree) RightNodeIsExist(root *Tree) (bool,*Tree) {
	if root.right==nil{
		return false,nil
	}else {
		return true,root.right
	}
}
//判断是否为叶子结点
func (tree *Tree) IsLeaf(node *Tree) bool {
	if node==nil {
		return false
	}
	if node.left==nil && node.right==nil {
		return true
	}
	return false
}
//打印当前节点的值
func (tree *Tree) PrintNodeValue(root *Tree) (nodevalue interface{}) {
	nodevalue=root.value
	return nodevalue
}
//树的深度
func (tree *Tree) TreeDepth(root *Tree) int {
	if root==nil {
		return 0
	}
	nleft:=tree.TreeDepth(root.left)
	nright:=tree.TreeDepth(root.right)
	if nleft>=nright{
		return nleft+1
	}else {
		return nright+1
	}
}
//判断该树是否为平衡二叉树
func (tree *Tree) IsBalanced1(root *Tree) bool {
	if root==nil{
		return true
	}
	left := tree.TreeDepth(root.left)
	right := tree.TreeDepth(root.right)
	differ:=left-right
	if differ < -1 || differ > 1 {
		return false
	}
	return tree.IsBalanced1(root.left)&&tree.IsBalanced1(root.right)
}
func (tree *Tree) IsBalanced2(root *Tree,depth *int) bool  {
	if root==nil {
		*depth=0
		return true
	}
	var left,right int
	if tree.IsBalanced2(root.left,&left)&&tree.IsBalanced2(root.right,&right) {
		differ := left-right
		if differ<=1&&differ>=-1 {
			if left>=right {
				*depth=left+1
			}else {
				*depth=right+1
			}
			return true
		}
	}
	return false
}
//判断该数组是否为二叉搜索树的后序遍历
func (tree *Tree) VerifySquenceOfBST(nums []int,front,end int) bool {
	if len(nums)<=0 {
		return false
	}
	if front>=end{
		return true
	}
	var i int
	for i=front;i<end ;i++  {
		if nums[i]>nums[end] {
			break;
		}
	}
	mid:=i-1
	for ;i<end ;i++  {
		if nums[i]<nums[end] {
			return false
		}
	}
	return tree.VerifySquenceOfBST(nums,front,mid)&&tree.VerifySquenceOfBST(nums,mid+1,end-1)
}
//判断两棵树是否有相同的值,并将相同的值入队
func (tree *Tree) Getsamenode(root1 ,root2 *Tree) queue.Queue{
	if root1 != nil {
		if root1.value==root2.value {
		que.EnQueue(root1)
		}
		tree.Getsamenode(root1.left,root2)
		tree.Getsamenode(root1.right,root2)
	}
		return que
}
//判断两棵树的结构是否相同
func (tree *Tree) DoesSameStruct(root1,root2 *Tree) bool {
	if root1!=nil&&root2!=nil {
		if root1.value==root2.value {
			return tree.DoesSameStruct(root1.left,root2.left)&&tree.DoesSameStruct(root1.right,root2.right)
		}else {
			return false
		}
	}else {
		if root2==nil || root2.value==nil{
			return true
		}else {
			return false
		}
	}
}
//二叉树的镜像
func (tree *Tree) MirrorRecursively(root *Tree)  {
	if root==nil||root.value==nil {
		return
	}
	if root.left==nil&&root.right==nil {
		return
	}
	que.EnQueue(root)
	defer que.Clear()
	for ! que.IsEmpty(){
		node := interface{}(que.DeQueue()).(*Tree)
		if node.left==nil {
			node.left=new(Tree)
		}
		if node.right==nil {
			node.right=new(Tree)
		}//初始化新节点
		node.left,node.right=node.right,node.left
		if !tree.IsLeaf(node.left) {
			que.EnQueue(node.left)
		}
		if !tree.IsLeaf(node.right) {
			que.EnQueue(node.right)
		}
	}
}
//对称树判断
//利用前序遍历和对称前序遍历判断该树是否为对称树，当前序遍历和对称前序遍历相同时为对称树
func (tree *Tree) IsSymmetrical(root *Tree) bool {
	return tree.isSymmetrical(root,root)
}
func (tree *Tree) isSymmetrical(root1,root2 *Tree) bool {
	if root1==nil && root2==nil {
		return true
	}
	if root1==nil || root2==nil {
		return false
	}
	if root1.value!=root2.value {
		return false
	}
	return tree.isSymmetrical(root1.left,root2.right)&&tree.isSymmetrical(root1.right,root2.left)
}
//打印出二叉树中和为n的路径
func (tree *Tree) FindPath(root *Tree,n,currentSum int)  {
	if root==nil {
		return
	}
	sta.Push(root)

	currentSum=currentSum+root.value.(int)
	if n==currentSum {
		nums :=sta.TraceStack()
		for _,value := range nums{
			fmt.Printf("%v ",value.(*Tree).value)
		}
		fmt.Print("\n")
		sta.Pop()
		return
	}
	if root.left!=nil {
		tree.FindPath(root.left,n,currentSum)
	}
	if root.right!=nil{
		tree.FindPath(root.right,n,currentSum)
	}
	sta.Pop()
}
//将二叉搜索树转变成有序的双向链表
func (tree *Tree) TreeConvertList(root *Tree,list *list.LinkedList) *list.LinkedList {
	if root != nil {
		list=tree.TreeConvertList(root.left,list)
		temp := ls.InitNode(root.value)
		temp.Pre=list
		list.Next=temp
		list=temp
		list=tree.TreeConvertList(root.right,list)
	}
	return list
}
//利用前序遍历序列化二叉树
func (tree *Tree) Serialize(root *Tree,qu **queue.Queue)  {
	q:= *qu
	if root==nil {
		q.EnQueue("$")
		return
	}
	q.EnQueue(root.value)
	tree.Serialize(root.left,&q)
	tree.Serialize(root.right,&q)
}
//利用前序遍历反序列化二叉树
func (tree *Tree) DeSerialize(qu *queue.Queue) *Tree {
	var temp *Tree
	if !qu.IsEmpty() {
		num:= qu.DeQueue()
		if num=="$" {
			temp=nil
		}else {
			temp= initnode(num)
			temp.left=tree.DeSerialize(qu)
			temp.right=tree.DeSerialize(qu)
		}
	}
	return temp
}

func (tree *Tree) MaxTreeLength(root *Tree) int {
	if root==nil{
		return 0
	}
	//fmt.Println(root.value)
	var length1,length2 int
	if root!=nil{
		length1 = tree.MaxTreeLength(root.left)
		length2 = tree.MaxTreeLength(root.right)
	}
	if length1>length2 {
		return length1+1
	}else {
		return length2+1
	}
}



