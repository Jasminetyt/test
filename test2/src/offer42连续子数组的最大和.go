package main

import "fmt"

var flag=true
func main() {
	var nums=[]int{-1,-1,-2,3,10,-4,7,2,-5}
	max := findGreatestSumOfSubArray(nums)
	if flag {
		fmt.Println("最大值为 ",max)
	}else {
		fmt.Println("数组输入错误")
	}
	dp := findGreatestSumOfSubArray2(nums)
	for _,value := range dp{
		fmt.Printf("%d ",value)
	}
	fmt.Println("")
	findGreatestSumOfSubArray3(nums)
}
func findGreatestSumOfSubArray(nums []int) int {
	if len(nums)==0 {
		flag=false
		return 0
	}
	var max,sum=0,0
	for i:=0;i<len(nums);i++  {
		if sum<0 {
			sum=nums[i]
		}else {
			sum=sum+nums[i]
		}
		if sum>max {
			max=sum
		}
	}
	return max
}
func findGreatestSumOfSubArray2(nums []int) []int {
	var dp=[]int{nums[0]}
	var sum=0
	for i:=0;i<len(nums);i++  {
		sum=sum+nums[i]
		if i>0{
			if sum<dp[i-1] {
				dp=append(dp,dp[i-1])
				}else {
					dp=append(dp,sum)
					}
		}
		if sum<0 {
			sum=0
		}
	}
	return dp
}
func findGreatestSumOfSubArray3(nums []int)  {
	var max,sum=0,0
	var result,temp []int
	for i:=0;i<len(nums);i++{
		sum=sum+nums[i]
		temp=append(temp,nums[i])
		if sum<0{
			sum=0
			temp=temp[:0]
		}
		if sum>max{
			max=sum
			result=temp
		}
	}
	fmt.Println(max,result)
}
