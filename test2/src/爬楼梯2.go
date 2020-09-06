package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	var dp = make([]string,4)
	dp[0],dp[1],dp[2],dp[3]="0","1","1","2"
	for i:=4;i<=n;i++{
		dp=append(dp,sum(dp[i-1],dp[i-3]))
	}
	fmt.Println(dp[len(dp)-1])
}
func sum(str1,str2 string) (str string) {
	chars1 := []byte(str1)
	chars2 := []byte(str2)
	var i,j=len(chars1)-1,len(chars2)-1
	var temp,flag=0,0
	for ;i>=0&&j>=0 ;  {
		num1:=int(chars1[i]-'0')
		num2:=int(chars2[j]-'0')
		temp=flag+num1+num2
		if temp>9 {
			temp=temp-10
			flag=1
		}else {
			flag=0
		}
		str=strconv.Itoa(temp)+str
		i--
		j--
	}
	if i>=0 {
		temp=flag+int(chars1[i]-'0')
		if temp>9 {
			temp=temp-10
			flag=1
		}else {
			flag=0
		}
		str=strconv.Itoa(temp)+str
		i--
	}
	if j>=0 {
		temp=flag+int(chars2[j]-'0')
		if temp>9 {
			temp=temp-10
			flag=1
		}else {
			flag=0
		}
		str=strconv.Itoa(temp)+str
		j--
	}
	if flag>0 {
		str=strconv.Itoa(flag)+str
	}
	return str
}
