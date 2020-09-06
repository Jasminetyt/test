package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var scan = bufio.NewScanner(os.Stdin)
	var str string
	for scan.Scan(){
		str=scan.Text()
		break
	}
	str=str[1:len(str)-1]
	strs:=strings.Split(str,",")
	var num=make([]int,len(strs))
	for i:=0;i<len(num);i++{
		temp,_:=strconv.Atoi(strs[i])
		num[i]=temp
	}
	var dp=make([]int,n+1)
	dp[0]=1
	for i:=0;i<len(num);i++{
		for j:=0;j<=n;j++{
			if j>=num[i]{
				dp[j]=dp[j]+dp[j-num[i]]
			}
		}
	}
	fmt.Println(dp[n])
}
