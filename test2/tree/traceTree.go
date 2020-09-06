package tree

import "fmt"

//前序遍历，递归
func (tree *Tree) PreOrderTrace(root *Tree)  {
	if root != nil {
		fmt.Println(root.value)
		tree.PreOrderTrace(root.left)
		tree.PreOrderTrace(root.right)
	}
}
//前序遍历，非递归
func (tree *Tree) PreOrderTrace2(root *Tree)  {
	if root==nil {
		return
	}
	tree.pushleftAndprint(root)
	for ! sta.IsEmpty()  {
		node := interface{}(sta.Pop()).(*Tree)
		if node.right !=nil {
			tree.pushleftAndprint(node.right)
		}
	}
	sta.Clear()
}
//将当前结点的全部左结点压入栈并输出
func (tree *Tree) pushleftAndprint(root *Tree)  {
	for root!=nil  {
		fmt.Println(root.value)
		sta.Push(root)
		root=root.left
	}

}
//中序遍历，递归
func (tree *Tree) InOrderTrace(root *Tree)  {
	if root != nil {
		tree.InOrderTrace(root.left)
		fmt.Println(root.value)
		tree.InOrderTrace(root.right)
	}
}
//中序遍历，非递归
func (tree *Tree)InOrderTrace2(root *Tree)  {
	if root==nil {
		return
	}
	tree.pushleft(root)
	for ! sta.IsEmpty()  {
		node := interface{}(sta.Pop()).(*Tree)
		fmt.Println(node.value)
		if node.right!=nil {
			node=node.right
			tree.pushleft(node)
		}
	}
	sta.Clear()
}
//将当前结点的全部左结点压入栈中
func (tree *Tree) pushleft(root *Tree)  {
	sta.Push(root)
	for root.left!=nil{
		sta.Push(root.left)
		root=root.left
	}
}
//后续遍历，递归
func (tree *Tree) PostOrderTrace(root *Tree)  {
	if root != nil {
		tree.PostOrderTrace(root.left)
		tree.PostOrderTrace(root.right)
		fmt.Println(root.value)
	}
}
//后续遍历，非递归
func (tree *Tree) PostOrderTrace2(root *Tree)  {
	if root==nil {
		return
	}
	var value interface{}
	tree.pushleft(root)
	for !sta.IsEmpty(){
		node := interface{}(sta.Front()).(*Tree)
		if node.right==nil {
			fmt.Println(node.value)
			value=node.value
			sta.Pop()
		}else {
			if value==node.right.value {
				fmt.Println(node.value)
				value=node.value
				sta.Pop()
			}else {
				tree.pushleft(node.right)
			}
		}
	}
	sta.Clear()
}
//层次遍历
func (tree *Tree) LevelOrderTrace(root *Tree)  {
	if root==nil {
		return
	}
	quee.EnQueue(root)
	for ! quee.IsEmpty(){
		node := interface{}(quee.DeQueue()).(*Tree)
		fmt.Printf("%v ",node.value)
		if node.left!=nil {
			que.EnQueue(node.left)
		}
		if node.right!=nil {
			que.EnQueue(node.right)
		}
	}
	quee.Clear()
}
