package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums=[]int{2,4,11,5}
	fmt.Println(canPartition(nums))
}
func canPartition(nums []int) bool {
	var sum int
	for i:=0;i<len(nums);i++{
		sum=sum+nums[i]
	}
	if sum%2!=0{
		return false
	}
	temp:=sum/2
	sort.Ints(nums)
	var dp=make([]bool,temp+1)
	if nums[0]<=temp{
		dp[nums[0]]=true
	}
	for i:=1;i<len(nums);i++{
		for j:=temp-nums[i];j>=0;j--{
			if dp[j]==true{
				dp[j+nums[i]]=true
			}
		}
		//fmt.Println(dp)
	}
	if dp[temp]==true{
		return true
	}
	return false
}