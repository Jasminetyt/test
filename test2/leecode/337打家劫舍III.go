package main

func main() {
	//成树状结构
}
type TreeNode struct {
	    Val int
	    Left *TreeNode
        Right *TreeNode
}

func rob3(root *TreeNode) int {
	if root==nil{
		return 0
	}
	var include=root.Val
	var exclude=rob3(root.Left)+rob3(root.Right)
	if root.Left!=nil{
		include=include+rob3(root.Left.Left)
		include=include+rob3(root.Left.Right)
	}
	if root.Right!=nil{
		include=include+rob3(root.Right.Left)
		include=include+rob3(root.Right.Right)
	}
	if include>exclude{
		return include
	}else {
		return exclude
	}
}