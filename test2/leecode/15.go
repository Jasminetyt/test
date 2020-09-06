package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums=[]int{-1,0,1,2,-1,-1,-4}//-4,-1,-1,-1,0,1,2 [-1,0,1,2,-1,-4]
	fmt.Println(threeSum(nums))
	fmt.Println(threeSum2(nums,0))
}
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	if len(nums)>2&&nums [0]<=0&&nums[len(nums)-1]>=0{
		for i:=0;i<len(nums)-2;{
			first :=i+1
			last:=len(nums)-1
			for first<last{
				sum:=nums[first]+nums[last]+nums[i]
				if sum==0 {
					temp:=[]int{nums[first],nums[i],nums[last]}
					result=append(result,temp)
				}
				if sum>=0  {
					for first<last&&nums[last]==nums[last-1]{
						last--
					}
					last--
				}else {
					for first<last&&nums[first]==nums[first+1]{
						first++
					}
					first++
				}
			}
			for i+1<len(nums)&&nums[i]==nums[i+1]{//优化数组中出现多个零的情况
				i++
			}
			i++
		}
	}
	return result
}
func threeSum2(nums []int,n int) [][]int {
	var result [][]int
	sort.Ints(nums)
	//fmt.Println(nums)
	for i:=0;i<len(nums)-2;i++{
	//	fmt.Println(result,i)
		first,last:=i+1,len(nums)-1
		for first<last{
			sum:=nums[first]+nums[i]+nums[last]
			if sum==n{
				temp:=[]int{nums[i],nums[first],nums[last]}
				result=append(result,temp)
			}
			if sum<n{
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

	}
	return result
}
