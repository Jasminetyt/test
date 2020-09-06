package main

import "fmt"

func main() {
	var amount=11
	var coins=[]int{2,4,11,5}
	fmt.Println(change(amount,coins))
}
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
}
