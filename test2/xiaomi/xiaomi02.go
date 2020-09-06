package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var nums=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&nums[i])
	}
	//fmt.Println(nums)
	var min,max=0,0
	var result int

	for i:=1;i<len(nums);i++{
		//fmt.Println(nums[i],nums[min])
		if nums[i]<nums[i-1]{
			result=result+nums[max]-nums[min]
			//fmt.Println("result",result)
			min=i
			max=i
			//flag=true
		}else {
			max=i
		}
	}

	if nums[max]>nums[min]{
		result=result+nums[max]-nums[min]
	}
	//fmt.Println(min,max)
	fmt.Println(result)
}
