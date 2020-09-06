package main

import "fmt"

func main() {
	var nums=[][]int{{1,10,3,8},{12,2,9,6},{5,7,4,11},{3,7,16,5}}
	if len(nums)==0 || nums==nil {
		return
	}
	fmt.Println(getMaxValue_solution1(nums))
	fmt.Println(getMaxValue_solution2(nums))
	getMaxValue_solution3(nums)
	getMaxValue_solution4(nums)
}
func getMaxValue_solution1(nums [][]int) int {
	var maxValues=make([][]int,len(nums))
	for i:=0;i<len(nums);i++{
		maxValues[i]=make([]int,len(nums[i]))
		for j:=0;j<len(nums[i]);j++{
			left := 0
			up :=0
			if i>0 {
				up=maxValues[i-1][j]
			}
			if j>0 {
				left=maxValues[i][j-1]
			}
			if up>left {
				maxValues[i][j]=up+nums[i][j]
			}else {
				maxValues[i][j]=left+nums[i][j]
			}
		}
	}
	return maxValues[len(nums)-1][len(nums[0])-1]
}
func getMaxValue_solution2(nums [][]int) int {
	maxValues := make([]int,len(nums[0]))
	for i:=0;i<len(nums);i++{
		for j:=0;j<len(nums) ;j++  {
			left :=0
			up := 0
			if i>0 {
				left=maxValues[j]
			}
			if j>0 {
				up=maxValues[j-1]
			}
			if left>up {
				maxValues[j]=nums[i][j]+left
			}else {
				maxValues[j]=nums[i][j]+up
			}
		}
	}
	return maxValues[len(nums[0])-1]
}

func getMaxValue_solution3(nums [][]int){
	var dp=make([][]int,len(nums))
	for i:=0;i<len(nums);i++{
		dp[i]=make([]int,len(nums[0]))
	}
	dp[0][0]=nums[0][0]
	for i:=1;i<len(nums[0]);i++{
		dp[0][i]=dp[0][i-1]+nums[0][i]
	}
	for i:=1;i<len(nums);i++{
		dp[i][0]=dp[i-1][0]+nums[i][0]
	}
	for i:=1;i<len(nums);i++{
		for j:=1;j<len(nums[0]);j++{
			if dp[i-1][j]>dp[i][j-1]{
				dp[i][j]=dp[i-1][j]+nums[i][j]
			}else {
				dp[i][j]=dp[i][j-1]+nums[i][j]
			}
		}
	}
	fmt.Println(dp[len(nums)-1][len(nums[0])-1])
}
func getMaxValue_solution4(nums [][]int){
	var dp=make([]int,len(nums[0]))
	for i:=0;i<len(nums);i++{
		dp[0]=nums[i][0]+dp[0]
		for j:=1;j<len(nums[0]);j++{
			if dp[j]>dp[j-1]{
				dp[j]=dp[j]+nums[i][j]
			}else {
				dp[j]=dp[j-1]+nums[i][j]
			}
		}
	}
	fmt.Println(dp[len(nums[0])-1])
}