package tree

//初始化树的节点
func initnode(num interface{}) *Tree {
	if num==nil {
		return nil
	}
	if num==nil{
		return nil
	}
	node:=new(Tree)
	node.value=num
	node.left=nil
	node.right=nil
	return node
}
//层次构建树
func (tree *Tree) InitTree(nums []interface{},root *Tree) *Tree {
	var i=0
	for i<len(nums)  {
		if que.IsEmpty() {
			root=initnode(nums[i])
			que.EnQueue(root)
			i++
		}else {
			node := interface{}(que.DeQueue()).(*Tree)
			nodeleft := initnode(nums[i])
			que.EnQueue(nodeleft)
			var noderight *Tree
			if i+1<len(nums) {
				i++
				noderight= initnode(nums[i])
				que.EnQueue(noderight)
			}
			if node!=nil{
				node.left=nodeleft
				node.right=noderight
			}
			i++
		}
	}
	quee.Clear()
	return root
}
//构建二叉搜索树
func (tree *Tree) InitBinaryNodeTree(nums []interface{},root *Tree) *Tree {
	if len(nums)<=0 {
		return nil
	}
	for i:=0;i<len(nums);i++{
		root=tree.insertBinaryNode(root,nums[i])
	}
	return root
}
func (tree *Tree) insertBinaryNode(node *Tree,num interface{}) *Tree {
	if node==nil||node.value==nil {
		node = initnode(num)
		return node
	}
	if num.(int)<=node.value.(int) {
		node.left=tree.insertBinaryNode(node.left,num)
	}else {
		node.right=tree.insertBinaryNode(node.right,num)
	}
	return node
}
//利用中序遍历和前序遍历创建二叉树
func (tree *Tree) CreateTreeByPreOrderAndInOrder(nums1,nums2 []interface{},preStart,preEnd,inStart,inEnd int)*Tree{
	if preStart>preEnd || inStart>inEnd{
		return nil
	}
	//node:= initnode(nums1[preStart])
	node:=new(Tree)
	node.value=nums1[preStart]
	for i:=inStart;i<=inEnd;i++{
		if nums2[i]==nums1[preStart]{
			node.left=tree.CreateTreeByPreOrderAndInOrder(nums1,nums2,preStart+1,preStart+i-inStart,inStart,i-1)
			node.right=tree.CreateTreeByPreOrderAndInOrder(nums1,nums2,preStart+i-inStart+1,preEnd,i+1,inEnd)
			break
		}
	}
	return node
}

