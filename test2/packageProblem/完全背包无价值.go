package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var coins=[]int{1,2,5}
	var amout=11
	fmt.Println(coinChange(coins,amout))
	fmt.Println(change(amout,coins))
}
func coinChange(coins []int, amount int) int {
	var dp=make([]int,amount+1)
	sort.Ints(coins)
	for i:=0;i<=amount;i++{
		dp[i]=math.MaxInt32
	}

	dp[0]=0
	for i:=0;i<len(coins);i++{
		for j:=coins[i];j<=amount;j++{
			min := dp[j]
			if dp[j-coins[i]]<min{
				min=dp[j-coins[i]]
				dp[j]=min+1
			}
		}
	}

	if dp[amount]!=math.MaxInt32{
		return dp[amount]
	}else {
		return -1
	}

} //输出最少次数
func change(amount int, coins []int) int {
	var dp=make([]int,amount+1)
	dp[0]=1
	for i:=0;i<len(coins);i++{
		for j:=0;j<=amount;j++{
			if j>=coins[i]{
				dp[j]=dp[j]+dp[j-coins[i]]
			}
		}
	}
	// 在上一钟零钱状态的基础上增大
	// 例如对于总额5，当只有面额为1的零钱时，只有一种可能 5x1
	// 当加了面额为2的零钱时，除了原来的那一种可能外
	// 还加上了组合了两块钱的情况，而总额为5是在总额为3的基础上加上两块钱来的
	// 所以就加上此时总额为3的所有组合情况
	fmt.Println(dp)
	return dp[amount]
} //输出可以组合的次数

