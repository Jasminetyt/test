package main

import "fmt"

func main() {
	var N ,W int
	fmt.Scan(&N)//数量
	fmt.Scan(&W)//承重
	var m=make([]int,N)//重量
	var v=make([]int,N)//价格
	for i:=0;i<N;i++{
		fmt.Scan(&m[i])
	}
	for i:=0;i<N;i++{
		fmt.Scan(&v[i])
	}
	var dp=make([][]int,N)
	for i:=0;i<N;i++{
		dp[i]=make([]int,W+1)
	}
	for i:=0;i<=W;i++{
		if m[0]<=i{
			dp[0][i]=v[0]
		}else {
			dp[0][i]=0
		}
	}
	//fmt.Println(dp)
	//fmt.Println(dp[1][0])
	for i:=1;i<N;i++{
		for j:=0;j<=W;j++{
		//	fmt.Println(i,j)
			//fmt.Println(dp[i][j])
			dp[i][j]=dp[i-1][j]
			//fmt.Println(dp[i][j])
			if m[i]<=j{
				if dp[i][j]<v[i]+dp[i-1][j-m[i]]{
					dp[i][j]=v[i]+dp[i-1][j-m[i]]
				}
			}
			//fmt.Println(dp)
		}
	}
	fmt.Println(dp[N-1][W])
}
