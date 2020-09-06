package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var nums=[]int{0,1,2}//-4,-1,1,2
	target :=0
	fmt.Println(threeSumClosest(nums,target))

}
func threeSumClosest(nums []int, target int) int {
	var count int
	sort.Ints(nums)
	if len(nums)>2 {
		count=nums[0]+nums[1]+nums[2]
		result:=math.Abs(float64(nums[0]+nums[1]+nums[2]-target))
		fmt.Println("result=",result)
		outer:
		for i:=0;i<len(nums)-2;{
			first := i+1
			last := len(nums)-1
			for first<last{
				sum := nums[i]+nums[first]+nums[last]
				temp:=math.Abs(float64(target-sum))
				fmt.Println("sum=",sum,"temp=",temp)
				if temp<float64(result) {
					result=temp
					count=sum
				}
				if sum==target {
					count=sum
					break outer
				}
				if sum<target {
					for first<last&&nums[first]==nums[first+1]{
						first++
					}
					first++
				}else {
						for first<last&&nums[last]==nums[last-1]{
							last--
						}
						last--
				}
			}
			for i+1<len(nums)&&nums[i]==nums[i+1]{
				i++
			}
			i++
		}
	}
	return count
}
