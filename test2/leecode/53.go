package main

import "fmt"

func main() {
	var nums=[]int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	var max=nums[0]
	var sum int
	for i:=0;i<len(nums);i++{
		sum=sum+nums[i]
		if sum>max{
			max=sum
		}
		if sum<0{
			sum=0
		}
	}
	return max
}
