package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	var x,y int
	var nums []int
	for i:=0;i<n;i++{
		fmt.Scan(&x,&y)
		for i:=0;i<x;i++{
			nums=append(nums,y)
		}
	}
	sort.Ints(nums)
	var i,j=0,len(nums)-1
	max:=nums[i]+nums[j]
	for i<j{
		if nums[i]+nums[j]>=max{
			max=nums[i]+nums[j]
		}else {
			break
		}
		i++
		j--
	}
	fmt.Println(max)

}
