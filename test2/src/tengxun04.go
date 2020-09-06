package main

import (
	"fmt"
	"sort"
)

func main() {
	var n,k int
	fmt.Scan(&n,&k)
	var nums=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&nums[i])
	}
	sort.Ints(nums)
	var index int
	for index=0;index<len(nums);index++{
		if nums[index]>0 {
			break
		}
	}
	var temp int
	for i:=0;i<k;i++{
		if index<len(nums){
			temp=nums[index]
			fmt.Println(temp)
		}else {
			fmt.Println(0)
		}

		for j:=index;j<len(nums);j++{
			nums[j]=nums[j]-temp
		}
		for j:=index;j<len(nums);j++{
			if nums[j]<=0 {
				index++
			}else {
					break
			}
		}
	}
}
