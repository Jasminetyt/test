package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var dp=make([]int,n+1)
	dp[0]=0
	dp[1]=1
	dp[2]=2
	for i:=3;i<len(dp);i++{
		dp[i]=dp[i-1]+dp[i-2]
	}
	fmt.Println(dp[n])
}
