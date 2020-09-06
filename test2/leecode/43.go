package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(multiply1("123",'6',0))
	fmt.Println(sum("123","987"))
	fmt.Println(multiply("123","456"))
}

func multiply(num1 string, num2 string) string {
	if num1=="0" || num2=="0"{
		return "0"
	}
	var i,count int
	var temp string
	var temp2="0"
	for i=len(num2)-1;i>=0;i--{
		temp=multiply1(num1,num2[i],count)
		temp2=sum(temp,temp2)
		count++
	}
	return temp2
}
func multiply1(num1 string,num2 byte,count int) string {
	var i,temp1 int
	var str,temp2 string
	for i=len(num1)-1;i>=0;i--{
		temp1,temp2=multiply2(num1[i],num2,temp1)
		str=temp2+str
	}
	if temp1>0 {
		str=strconv.Itoa(temp1)+str
	}
	for i=0;i<count;i++{
		str=str+"0"
	}
	return str
}

func multiply2(n1,n2 byte,temp int) (temp1 int,temp2 string) { //temp1为进位，temp2为当前值
	num1 := int(n1-'0')
	num2 := int(n2-'0')
	result := num1*num2+temp
	temp1=result/10
	temp2=strconv.Itoa(result%10)
	return temp1,temp2
}

func sum(num1,num2 string) string {
	var i,j,temp,sum int
	var str string
	for i,j=len(num1)-1,len(num2)-1;i>=0&&j>=0;{
		sum=int(num1[i]-'0')+int(num2[j]-'0')+temp
		temp=sum/10
		str=strconv.Itoa(sum%10)+str
		i--
		j--
	}
	for i>=0{
		sum=int(num1[i]-'0')+temp
		temp=sum/10
		str=strconv.Itoa(sum%10)+str
		i--
	}
	for j>=0{
		sum=int(num2[j]-'0')+temp
		temp=sum/10
		str=strconv.Itoa(sum%10)+str
		j--
	}
	if temp>0{
		str=strconv.Itoa(temp)+str
	}
	return str
}