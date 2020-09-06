package main

import (
	"bufio"
	list2 "container/list"
	"fmt"
	"os"
	stack2 "test2/stack"
)

func main() {
	var(
		str1 []string
		str2 []string
		st =new(stack2.Stack)
	)
	scan := bufio.NewScanner(os.Stdin)
	scan.Split( bufio.ScanWords) //去掉空格
	fmt.Println("压入栈的顺序")
	for scan.Scan() {
		if scan.Text()=="end" {
			break
		}
		str1=append(str1,scan.Text())
	}
	fmt.Println("弹出栈的顺序")
	for scan.Scan() {
		if scan.Text()=="end" {
			break
		}
		str2=append(str2,scan.Text())
	}
	if len(str1)==0 || len(str2)==0 {
		fmt.Println("输入为空")
		return
	}
	for i,j := 0,0; ;{
		if st.Front()==nil ||st.Front().(string)!=str2[j] {
			if i>=len(str1) {
				break
			}
			fmt.Println(i)
			fmt.Println("push",str1[i])
			st.Push(str1[i])
			i++
		}else {
				fmt.Println("pop",st.Front().(string))
				st.Pop()
				j++
		}
	}
	if st.IsEmpty() {
		fmt.Println("st2为弹出顺序")
	}else {
		fmt.Println("st2不是弹出顺序")
	}
	fmt.Println("------------------")
	var num1=[]int{1,2,3,4,5}
	var num2=[]int{4,5,3,2,1}
	fmt.Println(isPopOrder(num1,num2))
}
func isPopOrder(num1,num2 []int) bool {
	var list=list2.New()
	var i,j int
	for i,j=0,0;i<len(num1);{
		if num2[j]==num1[i]{
			fmt.Println("值相同不压入栈中",num1[i])
			i++
			j++
		}else if list.Len()>0&&num2[j]==list.Back().Value.(int){
				fmt.Println("栈顶元素与当前值相同，取出栈顶元素",num2[j])
				j++
				node:=list.Back()
				list.Remove(node)
		}else {
			fmt.Println("压入栈",num1[i])
			list.PushBack(num1[i])
			i++
		}
	}
	if j<len(num2)&&list.Len()!=0{
		fmt.Println("栈中还有元素")
		for j<len(num2)&&list.Len()>0&&num2[j]==list.Back().Value.(int){
			fmt.Println("栈顶元素与当前值相同",num2[j])
			j++
			node:=list.Back()
			list.Remove(node)
		}
	}
	if list.Len()==0&&j==len(num2){
		return true
	}
	fmt.Println("栈顶元素与当前值不同")
	return false
}
