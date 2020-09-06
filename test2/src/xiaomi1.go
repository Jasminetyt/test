package main

import (
	"fmt"
)

//小明同学需要对一个长度为 N 的字符串进行处理，他需要按照要求执行若干步骤，每个步骤都均为下面 2 种操作中的一种，2 种操作如下：
//TYPE 1. 从字符串结尾开始算起，将第 X 个字符之前的字符移动到字符串末尾
//TYPE 2. 输出字符串索引为 X 的字符
//小明尝试了很久没能完成，你可以帮他解决这个问题吗？
//
//输入描述:
//第一行，包含两个整数，字符串的长度 N 和操作次数T；
//第二行为要操作的原始字符串；
//
//之后每行都是要执行的操作类型 TYPE 和操作中 X 的值，均为整数。
//
//输入范围：
//字符串长度 N：1 <= N <= 10000
//操作次数 T：1 <= T <= 10000
//操作类型 TYPE：1 <= TYPE<= 2
//变量 X：0 <= X < N
//
//输出描述:
//操作的执行结果
//
//输入例子1:
//6 2
//xiaomi
//1 2
//2 0
//
//输出例子1:
//m
func main() {
	var length,count int
	var str string
	fmt.Scan(&length,&count)
	fmt.Scan(&str)
	var type1, x int
	var chars=[]byte(str)
	var result []byte
	var flag1,flag2 =0,0
	for i:=0;i<count ;i++  {
		fmt.Scan(&type1,&x)
		if type1==1 {
			flag1=flag1+(len(chars)-x)
			if flag1>len(chars) {
				flag1=flag1-len(chars)
			}
		}
		if type1==2 {
			flag2=flag1+x
			if flag2>len(chars) {
				flag2=flag2-len(chars)
			}
			result=append(result,chars[flag2])
		}
	}
	for _,value := range result{
		fmt.Printf("%c\n",value)
	}
}
