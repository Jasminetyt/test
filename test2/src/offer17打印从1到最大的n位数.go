package main

import (
	"fmt"
	"strconv"
)

//考虑大数的问题，可使用字符串进行表示
//使用相加或者全排列进行解决
var chars []byte
var str []string
var count int
func main() {
	count=2
	printToMaxofNDigits(count)
	fmt.Println(str)
	var temp []string
	printToMaxofNDigits2(count,temp)
}

func printToMaxofNDigits(n int) {
	if n<=0 {
		return
	}
	for i:=0;i<10;i++ {
		chars=append(chars, byte(i+'0'))//注：i+'0'才能转化成对应数字的字符
		printToMaxofNDigits(n-1)
		changeToString(chars)
		chars=chars[:count-n]
	}
	return
}

func changeToString(chars []byte)  {
	if len(chars)!=count {
		return
	}
	for i:=0;i<len(chars);i++ {
		if chars[i]!='0' {
			str= append(str, string(chars[i:]))
			break
		}
	}
}

func printToMaxofNDigits2(n int,temp []string)  {
	if n==0{
		printToDigit(temp)
		return
	}
	for i:=0;i<=9;i++{
		temp=append(temp,strconv.Itoa(i))
		printToMaxofNDigits2(n-1,temp)
		temp=temp[0:len(temp)-1]
	}
}
func printToDigit(temp []string){
	var str string
	var i int
	for i=0;i<len(temp);i++{
		if temp[i]!="0"{
			break
		}
	}
	for j:=i;j<len(temp);j++{
		str=str+temp[j]
	}
	fmt.Printf("%s ",str)
}