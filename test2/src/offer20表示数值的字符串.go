package main

import (
	"fmt"
)

type listnode struct {
	value byte
	next *listnode
}
func main() {
	var str string
	fmt.Scan(&str)
	//head := new(listnode)
	//var head2 **listnode
	//head2=&head
	//initlist(str,head2)
	//fmt.Println(isNumericc(head))

	fmt.Println(match4([]byte(str),0))
}
func isNumericc(head *listnode) bool {
	if head==nil {
		return false
	}
	var head2 **listnode
	var flag=true
	//整数部分
	head2=&head
	fmt.Println("整数")
	if ! plusANDmul(head2) { //包含符号和数字
	return false
	}
	fmt.Printf("---------------- \n")
	if flag&&head!=nil && (head.value=='e' || head.value=='E' ){
		flag=false
		fmt.Println("整数的指数部分")
		head=head.next
		head2=&head
		if ! plusANDmul(head2) { //包含符号和数字
			return false
		}
	}
	fmt.Printf("----------------\n")
	if flag&&head!=nil && head.value=='.' {
		fmt.Println("小数")
		head=head.next
		head2=&head
		if ! isnumber(head2){
			return false
		}
	}
	fmt.Printf("----------------\n")
	if flag&&head!=nil && (head.value=='e' || head.value=='E' ) {
		fmt.Println("小数的指数部分")
		head=head.next
		head2=&head
		if ! plusANDmul(head2) { //包含符号和数字
			return false
		}
	}
	if head!=nil {
		fmt.Println("有其它非法字符")
		return false
	}

	return true
}
func plusANDmul(head2 **listnode) bool  {
	if (*head2)==nil {
		return false
	}
	if (*head2).value=='+' || (*head2).value=='-' {
		fmt.Println("+ and -")
		(*head2)=(*head2).next
		if !isnumber(&(*head2)) {
			return false
		}
	}else if !isnumber(&(*head2))&&(*head2).value!='.' {
		return false
	}
	return true
}

func isnumber(head **listnode) bool{
	fmt.Println("number 0-9")
	if *head==nil || !((*head).value >= '0' && (*head).value <= '9'){
		return false
	}
	for (*head)!=nil &&(*head).value >= '0' && (*head).value <= '9'{
		(*head)=(*head).next
	}
	return true
}

//头插法
func initlist(str string,head **listnode)  {
	chars := []byte(str)
	node := new(listnode)
	var temp *listnode
	node.value=chars[len(chars)-1]
	node.next=nil
	for i:=len(chars)-2;i>=0;i-- {
		temp=new(listnode)
		temp.value=chars[i]
		temp.next=node
		node=temp
	}
	*head=node
}

func match4(strs []byte ,x int) bool{
	var flag=true
	for i:=x;i<len(strs);{
		if strs[i]=='+'||strs[i]=='-'{
			i++
			temp:=i
			i=isnumber2(i,strs)
			if temp==i{
				return false
			}
		}else if strs[i]=='.'&&flag{
			i++
			temp:=i
			i=isnumber2(i,strs)
			if temp==i{
				return false
			}
			flag=false
		}else if strs[i]=='e'||strs[i]=='E'{
			i++
			temp:=i
			i=isnumber2(i,strs)
			if temp==i&&(i>=len(strs)||(strs[i]!='-'&&strs[i]!='+')){
				return false
			}
			flag=false
		}else if strs[i]>='0'&&strs[i]<='9'{
			i=isnumber2(i,strs)
		}else {
			return false
		}
	}
	return true
}
func isnumber2(x int,strs []byte) int {
	if x>=len(strs){
		return x
	}
	if strs[x]<'0'||strs[x]>'9'{
		return x
	}
	var i int
	for i=x;i<len(strs);{
		if strs[i]>='0'&&strs[i]<='9'{
			i++
		}else {
				break
		}
	}
	return i
}