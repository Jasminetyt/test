package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	var num1,num2 string
	var result=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&num1,&num2)
		result[i]=getcount2(num1,num2)
	}
	for i:=0;i<len(result);i++{
		fmt.Println(result[i])
	}
}
func getcount2(num1,num2 string) int {
	temp1:=[]byte(num1)
	temp2:=[]byte(num2)
	var mmap=make(map[int]int)
	for i:=0;i<6;i++{
		if temp1[i]!=temp2[i]{
			mmap[i]=1
		}
	}
	//fmt.Println(mmap)
	count:=len(mmap)
	if count2(mmap,temp1,temp2){
		return count
	}else {
		return -1
	}
}
func count2(mmap map[int]int,temp1,temp2 []byte) bool {
	//fmt.Println(mmap)
	var flag bool
	if len(mmap)==0{
		return true
	}
	for key,_:=range mmap{
		temp:=temp1[key]
		temp1[key]=temp2[key]
		if isprime(temp1){
			delete(mmap,key)
			flag=true
		}else {
			temp1[key]=temp
		}
	}
	if flag{
		return count2(mmap,temp1,temp2)
	}else {
		return false
	}
}
func isprime(num []byte) bool {
	//fmt.Println(string(num))
	var i int
	for i=0;i<len(num);i++{
		if num[i]!='0'{
			break
		}
	}
	temp,_:=strconv.Atoi(string(num[i:]))
	if temp<=3{
		return temp>1
	}
	for i=2;i<temp;i++{
		if temp%i==0{
			return false
		}
	}
	return true
}
