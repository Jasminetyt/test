package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	dp:=make([]int,n+1) //绳子长度为n时的，最大乘积
	dp[0],dp[1],dp[2],dp[3]=0,0,1,2
	if n<=3{
		fmt.Println(dp[n])
		return
	}
	dp[0],dp[1],dp[2],dp[3]=0,1,2,3
	for i:=4;i<=n;i++{
		max:=0
		for j:=1;j<=i/2;j++{
			if dp[i-j]*dp[j]>max{
				max=dp[i-j]*dp[j]
			}
		}
		dp[i]=max
	}

	fmt.Println(dp[n])
}
