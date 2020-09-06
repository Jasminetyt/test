package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var nums=make([]int,n)
	for i:=0;i<len(nums);i++{
		fmt.Scan(&nums[i])
	}
	var temp=make([]int,n)
	var count,max int
	for i:=0;i<len(nums);i++{
		count=0
		max=0
		for j:=i+1;j<len(nums);j++{
			if nums[j]>nums[i]{
				break;
			}else if nums[j]==nums[i]{
				count++
				break;
			}else {
				if nums[j]>max{
					max=nums[j]
					count++
				}
			}
		}
		temp[i]=count
	}
	//fmt.Println(temp)
	max=0
	var i int
	for i=0;i<len(temp);i++{
		if temp[i]>temp[max]{
			max=i
		}
	}
	fmt.Println(nums[max])
}
