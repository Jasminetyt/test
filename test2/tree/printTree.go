package tree

import "fmt"

//分行从上到下打印二叉树
func (tree * Tree) PrintTreeByLine(root *Tree)  {
	if root==nil {
		return
	}
	var nextlevel,toBePrinted=0,1
	quee.EnQueue(root)
	for !quee.IsEmpty(){
		temp := quee.Front().(*Tree)
		fmt.Printf("%v ",temp.value)
		if temp.left!=nil{
			quee.EnQueue(temp.left)
			nextlevel++
		}
		if temp.right!=nil {
			quee.EnQueue(temp.right)
			nextlevel++
		}
		quee.DeQueue()
		toBePrinted--
		if toBePrinted==0 {
			fmt.Print("\n")
			toBePrinted=nextlevel
			nextlevel=0
		}
	}
}
//之字行打印二叉树
func (tree *Tree) PrintTreeByZhi(root *Tree)  {
	if root == nil {
		return
	}
	sta.Stack1.Push(root)
	var (
		temp *Tree
	)
	for ! sta.Stack1.IsEmpty() || ! sta.Stack2.IsEmpty()  {
		for !sta.Stack1.IsEmpty(){
			temp=sta.Stack1.Front().(*Tree)
			fmt.Printf("%v ",temp.value)
			sta.Stack1.Pop()
			if temp.left!=nil {
				sta.Stack2.Push(temp.left)
			}
			if temp.right!=nil {
				sta.Stack2.Push(temp.right)
			}
		}
		fmt.Print("\n")
		for !sta.Stack2.IsEmpty()  {
			temp=sta.Stack2.Front().(*Tree)
			fmt.Printf("%v ",temp.value)
			sta.Stack2.Pop()
			if temp.right!=nil {
				sta.Stack1.Push(temp.right)
			}
			if temp.left!=nil {
				sta.Stack1.Push(temp.left)
			}
		}
		fmt.Print("\n")
	}
}
