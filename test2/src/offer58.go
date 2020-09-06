package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var chars []byte
	read := bufio.NewReader(os.Stdin)
	chars,_,_=read.ReadLine()
	reverseSentence(chars)
	fmt.Println(string(chars))
	var n int
	fmt.Scanln(&n)
	leftRotateString(chars,n)
}
func reverse(chars []byte,start,end int)  {
	for start<end{
		chars[start],chars[end]=chars[end],chars[start]
		start++
		end--
	}
}
//翻转单词顺序
func reverseSentence(chars []byte)  {
	if chars==nil || len(chars)<=0 {
		return
	}
	reverse(chars,0,len(chars)-1)
	 i,j := 0,0
	for ;j<len(chars);j++{
		if chars[j]==' ' {
			reverse(chars,i,j-1)
			i=j+1
		}
	}
	reverse(chars,i,j-1) //don't forget
}
//左旋转字符串
func leftRotateString(chars []byte,n int)  {
	if chars==nil || n <= 0 || len(chars)<=0 || n>len(chars){
		fmt.Println("输入非法输入")
		return
	}
	reverse(chars,0,n-1)
	reverse(chars,n,len(chars)-1)
	reverse(chars,0,len(chars)-1)
	fmt.Println(string(chars))
}
